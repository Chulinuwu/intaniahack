<script lang="ts">
    import { colorMap } from '$lib/utils/colorMap';
    import { iconMap } from '$lib/utils/iconMap';

    export let type = '';
    export let pic = '';
    export let description = '';
    export let money = 0;
    export let happiness = 0;
    export let knowledge = 0;
    export let relationship = 0;

    const stats = [
        { name: 'money', value: money },
        { name: 'happiness', value: happiness },
        { name: 'knowledge', value: knowledge },
        { name: 'relationship', value: relationship }
    ];
</script>

<div class="card border border-white rounded-lg shadow-lg bg-gray-800 text-white w-[66px] h-[98px]">
    <div class={`header p-0.5 ${colorMap[type]} rounded-t-lg flex justify-center items-center`}>
        <img src={iconMap[type]} alt="Money Icon" class="h-2" />
    </div>
    <div class="image-container">
        <img src={pic} alt="" class="w-[66px] h-[60px] object-cover border-t border-b border-white" />
    </div>

    <div class="flex flex-col items-center justify-center text-center">
        <div class="text-[8px] font-medium leading-none">{description}</div>
        
        <div class="flex justify-center items-center gap-0.5">
            {#each stats as stat}
            <div class={`flex items-center justify-center`}>
                <span class={`w-2 h-2 mr-0.5 rounded-full border border-white ${colorMap[stat.name] || 'bg-gray-400'}`}></span>
                <!-- <img src={colorMap[stat.name]} alt={`${stat.name} icon`} class="w-2 h-2 mr-0.5" /> -->
                <span class="font-light text-[8px]">{stat.value}</span>
            </div>
            {/each}
        </div>
    </div>
</div>
