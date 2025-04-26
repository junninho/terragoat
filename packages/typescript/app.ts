interface User {
    name: string;
    role: string;
}

function merge(target: any, source: any) {
    for (const key in source) {
        // Prototype pollution vulnerability
        target[key] = source[key];
    }
    return target;
}

const user: User = { name: "John", role: "user" };
const malicious = JSON.parse('{"__proto__": {"isAdmin": true}}');
const merged = merge(user, malicious);

console.log(({} as any).isAdmin); // true 