import { h } from 'preact';
import { useState } from 'preact/hooks.js';
import htm from 'htm';

const html = htm.bind(h);

const ws = new WebSocket("ws://localhost:8080/chat");

ws.onopen = e => console.log("connection open", e);

ws.onmessage = e => console.log("server sends a message: ", e.data);

export default function Counter({ start }) {
    const { value, increment, decrement } = useCounter(start);

    return html`
        <div>Counter: ${value}</div>
        <button onClick=${increment}>Increment</button>
        <button onClick=${decrement}>Decrement</button>
    `;
}

const useCounter = (start) => {
    const [value, setValue] = useState(start);

    return {
        value,
        increment() {
            setValue(value => {
                value += 1;
                ws.send(JSON.stringify({ value }));
                return value;
            });
        },
        decrement() {
            setValue(value => {
                value -= 1;
                ws.send(JSON.stringify({ value }));
                return value;
            });
        },
    };
};