/**
 * @author Matteo Salvatore Giancola -> Emperor42, ->matthew.giancola42@gmail.com
 * @argument target The name of the component in question
 * @argument dir Correctly formatted directory name/path in which you would find the template (i.e. must end with / to match correctly)
 * @async
 * @description Allows for a template to be found from a different html file in the same directory or a different if DIR is set, usable as web components once set
 * If using the VENI system directly this function will be called by the server on rendering to find any and all custom elements 
 * If this function fails the page will simply have the text for the given elements listed
 */
void async function fetchTemplate (dir, target) {
    //get the imported document in templates:
    var templates = document.createElement( "template" );
    var htmlFileName = target+"html";
    if(dir){
        htmlFileName = dir+htmlFileName;
    }
    templates.innerHTML = await ( await fetch( htmlFileName ) ).text();

    //create the custom element
    customElements.define(
        target,
        class extends HTMLElement {
            constructor() {
            super();
            const template = templates.content;
            const shadowRoot = this.attachShadow({ mode: "open" }).appendChild(
                template.cloneNode(true)
            );
            }
        }
    );
}