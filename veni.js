"use strict";

/**
 * veni - JavaScript based solution to find HTML web components
 * Automatically discovers and registers HTML web components for reuse
 * Can be fully embedded in a script tag, hosted in a static page
 * No web server required to work
 */

class Veni {
  constructor(options = {}) {
    this.options = {
      templatePath: options.templatePath || null,
      scanCurrentDocument: options.scanCurrentDocument !== false,
      autoRegister: options.autoRegister !== false,
      componentPrefix: options.componentPrefix || 'wc-',
      ...options
    };
    this.discoveredComponents = [];
    this.initialized = false;
  }

  /**
   * Initialize the component discovery system
   */
  init() {
    if (this.initialized) return;
    
    this.initialized = true;
    
    if (this.options.scanCurrentDocument) {
      this.discoverFromTemplate();
      this.discoverFromDocument();
    }
    
    return this;
  }

  /**
   * Discover components from template.html file if specified
   */
  discoverFromTemplate() {
    if (!this.options.templatePath) return;
    
    const template = this.loadTemplate(this.options.templatePath);
    if (template) {
      const components = this.parseComponentsFromHTML(template);
      this.registerComponents(components);
    }
  }

  /**
   * Discover components from the current document
   */
  discoverFromDocument() {
    const components = this.parseComponentsFromHTML(document.documentElement.outerHTML);
    this.registerComponents(components);
  }

  /**
   * Load template HTML from file (simulated for browser environment)
   */
  loadTemplate(templatePath) {
    // In a real implementation, this would fetch the template
    // For this example, we'll look for a template element
    const template = document.querySelector(`#${templatePath.replace('.html', '').replace('/', '')} template`);
    if (template) {
      return template.innerHTML;
    }
    
    // Simulate loading a template
    console.warn(`Template not found: ${templatePath}`);
    return null;
  }

  /**
   * Parse HTML to find and extract web components
   */
  parseComponentsFromHTML(html) {
    const components = [];
    const parser = new DOMParser();
    const doc = parser.parseFromString(html, 'text/html');
    
    // Find all custom elements
    const customElements = doc.querySelectorAll(':not(:defined)');
    
    customElements.forEach(el => {
      if (el.tagName && el.tagName.includes('-')) {
        const component = this.extractComponent(el);
        if (component && !components.some(c => c.name === component.name)) {
          components.push(component);
        }
      }
    });
    
    // Also look for template tags with component definitions
    const templates = doc.querySelectorAll('template');
    templates.forEach(template => {
      const component = this.extractComponentFromTemplate(template);
      if (component && !components.some(c => c.name === component.name)) {
        components.push(component);
      }
    });
    
    return components;
  }

  /**
   * Extract component from an element
   */
  extractComponent(element) {
    const tagName = element.tagName.toLowerCase();
    const componentName = tagName.replace(this.options.componentPrefix, '');
    
    return {
      name: componentName,
      tagName: tagName,
      html: element.outerHTML || element.innerHTML,
      attributes: this.extractAttributes(element),
      css: this.extractStyles(element)
    };
  }

  /**
   * Extract component from template element
   */
  extractComponentFromTemplate(template) {
    const content = template.innerHTML;
    
    // Look for custom element definition
    const shadowHost = template.querySelector('slot') || template;
    const componentName = shadowHost.tagName || 'wc-undefined';
    
    return {
      name: componentName.replace(this.options.componentPrefix, ''),
      tagName: componentName,
      html: content,
      isTemplate: true
    };
  }

  /**
   * Extract attributes from an element
   */
  extractAttributes(element) {
    const attrs = {};
    Array.from(element.attributes).forEach(attr => {
      attrs[attr.name] = attr.value;
    });
    return attrs;
  }

  /**
   * Extract styles from an element
   */
  extractStyles(element) {
    const styles = {};
    const style = element.getAttribute('style');
    if (style) {
      styles.inline = style;
    }
    return styles;
  }

  /**
   * Register discovered components
   */
  registerComponents(components) {
    components.forEach(component => {
      this.registerComponent(component);
    });
  }

  /**
   * Register a single component
   */
  registerComponent(component) {
    if (this.autoRegisterComponent(component)) {
      this.discoveredComponents.push(component);
    }
  }

  /**
   * Auto-register a component (custom element)
   */
  autoRegisterComponent(component) {
    if (!customElements.get(component.tagName) && this.options.autoRegister) {
      try {
        this.defineComponent(component);
        return true;
      } catch (error) {
        console.error(`Failed to register component ${component.name}:`, error);
        return false;
      }
    }
    return false;
  }

  /**
   * Define a custom web component
   */
  defineComponent(component) {
    const { name, tagName } = component;
    
    class CustomElement extends HTMLElement {
      constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        this.shadowRoot.innerHTML = component.html;
      }
    }
    
    customElements.define(tagName, CustomElement);
  }

  /**
   * Find all discovered components
   */
  getComponents() {
    return [...this.discoveredComponents];
  }

  /**
   * Find a specific component by name
   */
  getComponent(name) {
    return this.discoveredComponents.find(c => c.name === name);
  }

  /**
   * Get component count
   */
  getComponentCount() {
    return this.discoveredComponents.length;
  }
}

/**
 * Global instance and auto-initialization
 */
const veni = new Veni({
  scanCurrentDocument: true
});

// Auto-initialize when DOM is ready
if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', () => {
    veni.init();
  });
} else {
  veni.init();
}

// Export for module systems
if (typeof module !== 'undefined' && module.exports) {
  module.exports = { Veni };
}

if (typeof window !== 'undefined') {
  window.veni = { Veni, veni };
}