<script lang="ts">
    import { onMount } from 'svelte';

    let timeLeft = 20 * 60;
    let timer: ReturnType<typeof setInterval>;

    // ฟังก์ชันสำหรับแปลงวินาทีเป็น MM:SS
    function formatTime(seconds: number) {
        const minutes = Math.floor(seconds / 60).toString().padStart(2, '0');
        const secs = (seconds % 60).toString().padStart(2, '0');
        return `${minutes}:${secs}`;
    }

    onMount(() => {
        timer = setInterval(() => {
            if (timeLeft > 0) {
                timeLeft -= 1;
            } else {
                clearInterval(timer); // หยุดการนับถอยหลังเมื่อถึง 0
            }
        }, 1000);

        return () => clearInterval(timer); // ล้าง timer เมื่อ component ถูกทำลาย
    });
</script>

<div class="flex flex-col gap-0">
    <div class="text-xs leading-none">TIME LEFT</div>
    <div class="text-lg leading-none text-red-500">{formatTime(timeLeft)}</div>
</div>