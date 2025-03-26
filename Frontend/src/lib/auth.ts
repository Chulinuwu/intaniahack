import { writable } from 'svelte/store';
import { browser } from '$app/environment';

// สร้าง store สำหรับ token
export const token = writable<string | null>(null);

// ฟังก์ชันจัดการ token
export function setToken(newToken: string | null) {
    token.set(newToken);
    if (browser) {
        document.cookie = `token=${newToken}; path=/; max-age=604800`;
    }
}

export function getToken() {
    if (browser) {
        const cookies = document.cookie.split(';');
        const tokenCookie = cookies.find(c => c.trim().startsWith('token='));
        return tokenCookie ? tokenCookie.split('=')[1] : null;
    }
    let currentToken;
    token.subscribe(t => currentToken = t)();
    return currentToken;
}

export function removeToken() {
    token.set(null);
    if (browser) {
        document.cookie = 'token=; path=/; max-age=0';
    }
}

// ตั้งค่า initial value จาก cookie เมื่ออยู่ใน browser
if (browser) {
    const initialToken = getToken() ?? null;
    token.set(initialToken);
}