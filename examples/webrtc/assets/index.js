import { h, Component, render } from 'preact';
import htm from 'htm';
import Counter from './components/Counter.js';

const html = htm.bind(h);

function App(props) {
    return html`
        <h1>Hello ${props.name}!</h1>
        <${Counter} start="${0}"/>
    `;
}

render(html`<${App} name="Kristopher" />`, document.body);