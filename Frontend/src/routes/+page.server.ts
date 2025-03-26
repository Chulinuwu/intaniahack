// src/routes/+page.server.js
import { fail } from '@sveltejs/kit';

export const actions = {
    setToken: async ({ cookies, request }) => {
        const data = await request.formData();
        const token = data.get('token');
        if (typeof token !== 'string') {
            return fail(400, { message: 'Invalid token type' });
        }
        if (!token) {
            return fail(400, { message: 'No token provided' });
        }
        cookies.set('token', token, {
            path: '/',
            maxAge: 604800, // 7 วัน
            // secure: true, // เปิดใน production ที่มี HTTPS
            // httpOnly: true, // ป้องกัน XSS
            // sameSite: 'strict' // ป้องกัน CSRF
        });
        return { success: true };
    }
};