export default function auth(request: any): boolean {
    const token = request.headers.get("Authorization");
    return token === process.env.AUTH;
}
