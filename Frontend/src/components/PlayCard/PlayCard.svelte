<script lang="ts">
    import { colorMap } from '$lib/utils/colorMap';
    import { iconMap } from '$lib/utils/iconMap';
    import { iconMapColor } from '$lib/utils/iconMapColor';

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

<div class="card border border-white rounded-lg shadow-lg bg-gray-800 text-white w-[117px] h-[173px]">
    <div class={`header p-1 ${colorMap[type]} rounded-t-lg flex justify-center items-center`}>
        <img src={iconMap[type]} alt="Money Icon" class="h-3.5" />
    </div>
    <div class="image-container">
        <img src={pic} alt="" class="w-full h-[107px] object-cover border border-white" />
    </div>

    <div class="flex flex-col px-2 items-center justify-center text-center">
        <h2 class="text-xs font-medium">{description}</h2>
        <p class="text-[5px] text-gray-300">COST</p>
        
        <div class="flex justify-center gap-1">
            {#each stats as stat}
            <div class={`flex items-center px-1 rounded`}>
                <img src={iconMapColor[stat.name]} alt={`${stat.name} icon`} class="w-2 h-2 mr-1" />
                <span class="font-medium text-[10px]">{stat.value}</span>
            </div>
            {/each}
        </div>
    </div>
</div>
