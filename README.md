# VENI (Web Component Discovery System)

**MIT License © Matthew Salvatore Giancola**

## Overview

VENI is a JavaScript-based solution designed to discover and integrate HTML web components. It automatically identifies and adds web components to the Document object for reuse throughout web applications. No web server is required - it works with static HTML pages and can be fully embedded in script tags.

## Installation

### Prerequisites
- Modern web browser with JavaScript support
- No external server dependencies required

### Installation Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/emperor42/veni.git
   cd veni
   ```

2. Include VENI in your HTML:
   ```html
   <!-- Option 1: Direct script tag -->
   <script src="veni.js"></script>

   <!-- Option 2: Local file -->
   <script>
   // VENI will automatically discover and register components
   </script>
   ```

3. For Node.js development (optional):
   ```bash
   # Install Node.js build tools if needed
   npm install
   
   # Build for distribution (optional)
   npm run build
   ```

## Usage (Standalone)

### Basic Usage

```javascript
// Basic VENI usage
<script src="veni.js"></script>

// VENI will automatically discover components
// Components are available via window.VENI API

// Example: Use discovered components
const component = window.VENI.getComponent('my-component');
component.render();
```

### Advanced Usage

```javascript
// Initialize VENI with configuration
const veni = new VENI({
    autoDiscover: true,
    templateFile: 'components.html',
    customSelectors: ['.ven-component', '[data-venue]'],
    logLevel: 'info'
});

// Discover components from specific file
veni.discoverFromFile('components.html');

// Discover components from current HTML
veni.discoverFromCurrent();

// Get all discovered components
const components = veni.getComponents();

// Get component by name
const specificComponent = veni.getComponent('header-component');

// Register custom component
veni.registerComponent('custom-widget', {
    template: 'custom-template.html',
    styles: 'custom-widget.css'
});
```

### API Endpoints

| Method | Description |
|--------|-------------|
| `new VENI(options)` | Create new VENI instance |
| `veni.discoverFromFile(filename)` | Discover components from file |
| `veni.discoverFromCurrent()` | Discover components from current HTML |
| `veni.getComponents()` | Get all discovered components |
| `veni.getComponent(name)` | Get specific component |
| `veni.registerComponent(name, config)` | Register custom component |
| `veni.setAutoDiscover(enabled)` | Enable/disable auto-discovery |
| `veni.setLogLevel(level)` | Set log level |

### Example HTML Page

```html
<!DOCTYPE html>
<html>
<head>
    <title>VENI Component Demo</title>
    <!-- Include VENI -->
    <script src="venI.js"></script>
    
    <!-- Your content -->
    <!-- VENI will discover components from this file -->
    <!-- <my-component>My Custom Component</my-component> -->
</head>
<body>
    <!-- Discovered components will be available here -->
    <div id="content"></div>
    
    <!-- Use discovered components -->
    <script>
    // After VENI loads, components are available
    document.addEventListener('VENI-ready', function() {
        const headerComponent = window.VENI.getComponent('my-component');
        if (headerComponent) {
            headerComponent.render(document.getElementById('content'));
        }
    });
    </script>
</body>
</html>
```

## Integration with ATP

### Component Integration

VENI integrates with ATP to provide centralized component management:

```javascript
// VENI with ATP integration
const veni = new VENI({
    autoDiscover: true,
    templateEndpoint: 'https://api.example.com/atp/templates',
    componentEndpoint: 'https://api.example.com/atp/components',
    syncWithATP: true,
    onComponentsLoaded: function(components) {
        // Process ATP-integrated components
        components.forEach(component => {
            window.VENI.registerComponent(component.name, component.config);
        });
    }
});

// Synchronize with ATP
veni.syncWithATP().then(() => {
    console.log('Components synchronized with ATP');
    const atpComponents = veni.getComponents();
    console.log(`Loaded ${atpComponents.length} components from ATP`);
});
```

### Component Discovery

```javascript
// Discover components from ATP
const discoverFromATP = async () => {
    try {
        const response = await fetch('/atp/api/components');
        const components = await response.json();
        
        components.forEach(component => {
            window.VENI.registerComponent(component.name, {
                template: component.html,
                styles: component.css,
                scripts: component.js,
                attributes: component.attributes
            });
        });
        
        console.log(`Discovered ${components.length} components from ATP`);
    } catch (error) {
        console.error('Failed to discover components from ATP:', error);
    }
};
```

### Configuration Distribution

```javascript
// VENI configuration for ATP integration
const veniConfig = {
    autoDiscover: true,
    discoveryInterval: 30000, // 30 seconds
    templateFile: 'components.html',
    customSelectors: ['.venue-component', '[data-venue]'],
    apiEndpoint: '/atp/api/components',
    syncStrategy: 'merge', // merge, replace, append
    onSync: function(data) {
        console.log('Components synchronized:', data);
    },
    onError: function(error) {
        console.error('Component discovery error:', error);
    }
};
```

## Development Setup

### Local Development

```bash
# Test in browser
# Open browser and load:
# http://localhost:8080/venI.html
# (Load all VENI components)

# Or with Node.js
node -e "require('veni').test()"
```

### Testing

```javascript
// Basic VENI usage test
const VENI = window.VENI;

const testVeni = () => {
    // Test VENI initialization
    const veni = new VENI({
        autoDiscover: true,
        templateFile: 'components.html'
    });
    
    expect(veni).toBeDefined();
    expect(veni.autoDiscover).toBe(true);
    
    // Test component discovery
    veni.discoverFromCurrent();
    const components = veni.getComponents();
    expect(components).toBeDefined();
};

// Component registration test
const testComponentRegistration = () => {
    const veni = new VENI();
    
    veni.registerComponent('test-component', {
        template: '<div>Test Component</div>',
        styles: 'test-component.css'
    });
    
    const component = veni.getComponent('test-component');
    expect(component).toBeDefined();
    expect(component.name).toBe('test-component');
};
```

### Building

```bash
# Build for distribution
npm run build

# Output: dist/veni.js (optimized and bundled)

# Test in browser
# Open browser and load: dist/veni.js
```

## API Specifications

### High Maturity API (Event-driven)

#### Component Management
- `new VENI(options)` - Create new VENI instance
- `veni.discoverFromFile(filename)` - Discover components from file
- `veni.discoverFromCurrent()` - Discover components from current HTML
- `veni.getComponents()` - Get all discovered components
- `veni.getComponent(name)` - Get specific component
- `veni.registerComponent(name, config)` - Register custom component
- `veni.setAutoDiscover(enabled)` - Enable/disable auto-discovery
- `veni.setLogLevel(level)` - Set log level

#### Component Discovery
```javascript
// Discover components from file
veni.discoverFromFile('components.html');\n
// Discover components from current HTML
veni.discoverFromCurrent();

// Get all discovered components
const components = veni.getComponents();

// Get specific component
const component = veni.getComponent('header-component');

// Register custom component
veni.registerComponent('custom-widget', {
    template: 'custom-template.html',
    styles: 'custom-widget.css'
});
```

### VENI-specific Events

```javascript
// Component discovery event
veni.on('component-discovered', (event) => {
    const { component } = event;
    console.log(`Component discovered: ${component.name}`);
});

// Component registration event
veni.on('component-registered', (event) => {
    const { component } = event;
    console.log(`Component registered: ${component.name}`);
});

// Component discovery complete event
veni.on('discovery-complete', (event) => {
    const { componentCount } = event;
    console.log(`Discovery complete: ${componentCount} components found`);
});
```

## Security API

### Component Security
- `veni.setSecureMode(enabled)` - Enable secure mode
- `veni.setTrustedOrigins(origins)` - Set trusted origins
- `veni.validateComponent(component)` - Validate component security
- `veni.sanitizeComponent(component)` - Sanitize component

### Discovery Security
- `veni.setAutoDiscover(secure)` - Secure auto-discovery
- `veni.setDiscoveryTimeout(timeout)` - Set discovery timeout
- `veni.enableDiscoveryValidation(enabled)` - Enable discovery validation

## Integration API

### VENI Integration
- `veni.syncWithATP(config)` - Synchronize with ATP
- `veni.getVENIComponents()` - Get VENI components
- `veni.applyVENIComponents(components)` - Apply VENI components

### VENI Component Integration
```javascript
// Synchronize with VENI
veni.syncWithATP({
    endpoint: '/atp/api/components',
    strategy: 'merge',
    onSuccess: function(components) {
        console.log('Components synchronized:', components.length);
    }
});

// Get VENI components
const components = veni.getVENIComponents();

// Apply VENI components
veni.applyVENIComponents(components);
```

## Monitoring API

### Component Monitoring
- `veni.onComponentAdded(callback)` - Component added callback
- `veni.getComponentLogs()` - Get component logs
- `veni.getComponentMetrics()` - Get component metrics

### Discovery Monitoring
- `veni.onDiscoveryStart(callback)` - Discovery start callback
- `veni.onDiscoveryComplete(callback)` - Discovery complete callback
- `veni.getDiscoveryLogs()` - Get discovery logs

## Error Handling

### VENI Error Types
- `ComponentError` - Component discovery errors
- `SecurityError` - Security-related errors
- `ValidationError` - Input validation errors
- `IntegrationError` - Integration-related errors

### Error Response Format
```javascript
// VENI errors
class VENIError extends Error {
  constructor(message, code, details) {
    super(message);
    this.code = code;
    this.details = details;
    this.timestamp = new Date().toISOString();
  }
}
```

## Testing

### Unit Tests

```javascript
// Test VENI initialization
 test('VENI Initialization', () => {
   const veni = new VENI({
     autoDiscover: true,
     templateFile: 'components.html'
   });
   expect(veni).toBeDefined();
   expect(veni.autoDiscover).toBe(true);
 });

// Test component discovery
 test('Component Discovery', () => {
   const veni = new VENI({
     autoDiscover: true,
     templateFile: 'components.html'
   });
   
   veni.discoverFromCurrent();
   const components = veni.getComponents();
   expect(components).toBeDefined();
 });

// Test component registration
 test('Component Registration', () => {
   const veni = new VENI();
   
   veni.registerComponent('test-component', {
     template: '<div>Test Component</div>',
     styles: 'test-component.css'
   });
   
   const component = veni.getComponent('test-component');
   expect(component).toBeDefined();
   expect(component.name).toBe('test-component');
 });
```

### Integration Tests

```javascript
// Test VENI integration
 test('VENI-ATP Integration', () => {
   const veni = new VENI({
     autoDiscover: true,
     apiEndpoint: '/atp/api/components'
   });
   
   veni.syncWithATP().then(() => {
     const components = veni.getComponents();
     expect(components).toBeDefined();
   });
 });
```

## Performance Considerations

- **Memory Usage**: Monitor for large component sets
- **CPU Usage**: Optimize component discovery algorithms
- **Network I/O**: Cache frequently accessed components
- **Disk I/O**: Use efficient storage for component data
- **Concurrent Processing**: Support for concurrent component discovery

## Future Enhancements

- **Advanced Discovery**: AI-powered component discovery
- **Component Templates**: Component template generation
- **Component Analytics**: Component usage analytics
- **Advanced Integration**: Enhanced integration capabilities
- **Cloud Integration**: Integrate with cloud component services

## Troubleshooting

### Common Issues

1. **VENI not loading**
   ```javascript
   // Check VENI configuration
   const veni = new VENI({
     autoDiscover: true,
     templateFile: 'components.html'
   });
   
   // Test VENI initialization
   console.log('VENI initialized:', veni);
   ```

2. **Component discovery errors**
   ```javascript
   // Check component discovery
   veni.discoverFromCurrent();
   
   // Check component logs
   const logs = veni.getComponentLogs();
   console.log(logs);
   ```

3. **Event listener issues**
   ```javascript
   // Check event listener registration
   veni.on('component-discovered', (event) => {
     console.log('Component discovered:', event.component);
   });
   ```

### Debugging Commands

```javascript
// Enable debug logging
veni.setLogLevel('debug');

// Check component logs
const logs = veni.getComponentLogs();
console.log(logs);

// Monitor component discovery
veni.on('discovery-start', () => {
    console.log('Component discovery started');
});
```

## Conclusion

VENI provides an automated web component discovery and registration solution that simplifies the process of working with web components. It automatically discovers and registers components, making them available for use throughout a web application without complex setup processes.

Key benefits:

- **Automatic Discovery**: Automatic component discovery without manual configuration
- **WebComponent Integration**: Uses JavaScript standard WebComponent class
- **Flexible Templates**: Can work with specific template files or components in existing HTML
- **Embeddable Solution**: Can be fully embedded in script tags within static pages
- **No Server Requirement**: Operates independently without needing a web server
- **Secure Integration**: Secure integration with ATP platform
- **Production Ready**: Comprehensive error handling and monitoring

The VENI implementation is production-ready and can be easily integrated into web applications with comprehensive web component discovery and registration capabilities.

---

*Document Version: 1.0*
*Created: 2026-08-25*
*Last Updated: 2026-08-25*
*Status: Production Ready*

**License:** MIT License © Matthew Salvatore Giancola.