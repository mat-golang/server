import { h } from 'preact';
import { useState } from 'preact/hooks.js';
import htm from 'htm';

const html = htm.bind(h);

export default function Counter({ start }) {
    const [value, setValue] = useState(start);

    return html`
        <div>Counter: ${value}</div>
        <button onClick=${() => setValue(value + 1)}>Increment</button>
        <button onClick=${() => setValue(value - 1)}>Decrement</button>
    `;
}