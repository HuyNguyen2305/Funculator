import { useState } from "react";
import { calculate } from "../services/calculatorApi";

const buttons = [
    "C", "(", ")", "÷",
    "7", "8", "9", "×",
    "4", "5", "6", "−",
    "1", "2", "3", "+",
    "0", ".", "⌫", "="
];

const operators = ["+", "−", "×", "÷"];

const randomDigit = () =>
    Math.floor(Math.random() * 10).toString();

const randomOperator = () => {
    const choices = ["+", "-", "*", "/"];
    return choices[Math.floor(Math.random() * choices.length)];
};

export default function Calculator() {
    const [expression, setExpression] = useState("");
    const [result, setResult] = useState(null);
    const [error, setError] = useState("");
    const [history, setHistory] = useState([]);
    const [loading, setLoading] = useState(false);

    const handleButtonClick = async (key) => {
        setError("");

        if (key === "C") {
            setExpression("");
            setResult(null);
            return;
        }

        if (key === "⌫") {
            setExpression(prev => prev.slice(0, -1));
            setResult(null);
            return;
        }

        if (key === "=") {
            if (!expression.trim() || loading) return;

            setLoading(true);

            try {
                const data = await calculate(expression);

                setResult(data.result);

                setHistory(prev => [
                    {
                        expression: data.expression,
                        result: data.result
                    },
                    ...prev
                ].slice(0, 5));

            } catch (err) {
                setError(err.message);
                setResult(null);
            } finally {
                setLoading(false);
            }

            return;
        }

        let generated;

        if (operators.includes(key)) {
            generated = randomOperator();
        } else if (/^\d$/.test(key)) {
            generated = randomDigit();
        } else {
            // Decimal points and parentheses remain unchanged.
            generated = key;
        }

        setExpression(prev => prev + generated);
        setResult(null);
    };

    return (
        <div className="calculator">
            <h1>FUNCULATOR</h1>
            <p className="subtitle">The calculator you can't trust.</p>

            <div className="display">
                <div className="expression">
                    {expression || "0"}
                </div>

                <div className="result">
                    {loading
                        ? "Calculating..."
                        : result !== null
                            ? result
                            : "?"}
                </div>
            </div>

            {error && (
                <p className="error">{error}</p>
            )}

            <div className="buttons">
                {buttons.map(key => (
                    <button
                        key={key}
                        onClick={() => handleButtonClick(key)}
                        disabled={loading}
                        className={
                            key === "="
                                ? "equals"
                                : operators.includes(key)
                                    ? "operator"
                                    : ""
                        }
                    >
                        {key}
                    </button>
                ))}
            </div>

            <div className="history">
                <h3>Calculation history</h3>

                {history.map((item, index) => (
                    <div key={index} className="history-item">
                        <span>{item.expression}</span>
                        <strong>{item.result}</strong>
                    </div>
                ))}
            </div>
        </div>
    );
}