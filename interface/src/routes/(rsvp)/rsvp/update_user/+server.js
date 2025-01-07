export async function POST({ request }) {

    const household  = await request.json();

    console.log("Here is our request body")
    console.log(JSON.stringify(household));


    try {
        const response = await fetch(`http://localhost:8080/household/${household.id}`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(household),
        });

        if (response.ok) {
            return new Response('User updated successfully', { status: 201 });
        } else {
            const errorText = await response.text();
            return new Response(errorText, { status: response.status });
        }
    } catch (error) {
        console.error('Error communicating with the Go API:', error);
        return new Response('Internal server error', { status: 500 });
    }
}