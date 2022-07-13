import { h } from 'preact';
import { useState } from 'preact/hooks.js';
import htm from 'htm';

const html = htm.bind(h);

const ws = new WebSocket("ws://localhost:8080/chat");



export default function Counter({ start }) {
    const { value, onIncrement, onDecrement } = useCounter({ start, ws });


    ws.addEventListener("open", e => console.log("connected to the websocket "));

    return html`
        <div>Counter: ${value}</div>
        <button onClick=${onIncrement}>Increment</button>
        <button onClick=${onDecrement}>Decrement</button>
    `;
}

const useCounter = ({ start, ws }) => {
    const [value, setValue] = useState(start);

    ws.onopen = e => console.log("connection open", e);

    ws.onmessage = e => console.log("server sends a message: ", e.data);

    return {
        value,
        onIncrement() {
            setValue(value => {
                value += 1;
                ws.send(JSON.stringify({ value }));
                return value;
            });
        },
        onDecrement() {
            setValue(value => {
                value -= 1;
                ws.send(JSON.stringify({ value }));
                return value;
            });
        },
        onOpen() { },
        onMessage() { },
        onError() { },
    };
};