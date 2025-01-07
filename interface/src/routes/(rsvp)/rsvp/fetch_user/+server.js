export async function POST({ request }) {
    const { name } = await request.json();
    const encodedName = encodeURIComponent(name);

    try {
        const response = await fetch(`http://api:8080/household/${encodedName}`, {
            method: 'GET',
        });

        if (response.ok) {
            const data = await response.json();
            return new Response(JSON.stringify(data), { status: 200 });
        } else {
            return new Response('User not found', { status: response.status });
        }
    } catch (error) {
        console.error('Error communicating with the Go API:', error);
        return new Response('Internal server error', { status: 500 });
    }
}