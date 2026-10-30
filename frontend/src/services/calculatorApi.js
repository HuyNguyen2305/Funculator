const API_URL = "http://localhost:8080/api";

export async function calculate(expression) {
    const response = await fetch(`${API_URL}/calculate`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ expression }),
    });

    const data = await response.json();

    if (!response.ok) {
        throw new Error(data.error || "Calculation failed");
    }

    return data;
}