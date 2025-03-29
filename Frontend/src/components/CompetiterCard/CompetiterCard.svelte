<script lang="ts">
    import { colorMap } from '$lib/utils/colorMap';

    export let profileImage = '';
    export let playerName = '';
    export let money = 0;
    export let happiness = 0;
    export let knowledge = 0;
    export let relationship = 0;
    
    export let age0_12: { type: string, description: string }[] = [];
    export let age13_18: { type: string, description: string }[] = [];
    export let age19_22: { type: string, description: string }[] = [];
    export let age23_39: { type: string, description: string }[] = [];
    export let age40_59: { type: string, description: string }[] = [];
    export let age60_79: { type: string, description: string }[] = [];
    export let age80_100: { type: string, description: string }[] = [];

    let ageRanges = [
        { label: "0 - 12", data: age0_12 },
        { label: "13 - 18", data: age13_18 },
        { label: "19 - 22", data: age19_22 },
        { label: "23 - 39", data: age23_39 },
        { label: "40 - 59", data: age40_59 },
        { label: "60 - 79", data: age60_79 },
        { label: "80 - 100", data: age80_100 }
    ];
    
    let currentAgeIndex = 1;

    function changeAge(direction: number) {
        currentAgeIndex = Math.max(0, Math.min(ageRanges.length - 1, currentAgeIndex + direction));
    }
</script>

<div class="flex flex-col bg-gray-900 border border-white text-white rounded-lg p-2 gap-1 w-[116px] h-[217px]">
    <!-- Profile Section -->
    <div class="flex items-center">
        <img src={profileImage} alt="Profile" class="w-6 h-6 rounded-full border-2 border-white mr-3" />
        <h2 class="text-xs font-bold">{playerName}</h2>
    </div>

    <!-- Stats Section -->
    <div class="space-y-0 text-xs">
        <div class="flex justify-between items-center">
            <div class="flex items-center">
                <span class="w-2 h-2 bg-green-400 rounded-full border border-white mr-2"></span>
                <span>MONEY</span>
            </div>
            <span>{money}</span>
        </div>
        <div class="flex justify-between items-center">
            <div class="flex items-center">
                <span class="w-2 h-2 bg-yellow-400 border border-white rounded-full mr-2"></span>
                <span>HAPPINESS</span>
            </div>
            <span>{happiness}</span>
        </div>
        <div class="flex justify-between items-center">
            <div class="flex items-center">
                <span class="w-2 h-2 bg-purple-400 border border-white rounded-full mr-2"></span>
                <span>KNOWLEDGE</span>
            </div>
            <span>{knowledge}</span>
        </div>
        <div class="flex justify-between items-center">
            <div class="flex items-center">
                <span class="w-2 h-2 bg-blue-400 border border-white rounded-full mr-2"></span>
                <span>RELATIONSHIP</span>
            </div>
            <span>{relationship}</span>
        </div>
    </div>

    <!-- Age Navigation -->
    <div class="flex items-center justify-between">
        <button class="text-white" on:click={() => changeAge(-1)}>&#9665;</button>
        <div class="flex flex-col items-center leading-none">
            <div class="text-[12px]">{ageRanges[currentAgeIndex].label}</div>
            <div class="text-[8px]">age</div>
        </div>
        <button class="text-white" on:click={() => changeAge(1)}>&#9655;</button>
    </div>

    <!-- Age-Based Selections -->
    <div class="flex flex-col gap-1 text-xs">
        {#each ageRanges[currentAgeIndex].data as item}
            <div class="flex gap-1 items-start">
                <span class={`w-2 h-2 rounded-full border border-white flex-shrink-0 ${colorMap[item.type] || 'bg-gray-400'}`}></span>
                <p class="text-gray-400 text-[8px] leading-none flex-grow">{item.description}</p>
            </div>
        {/each}
    </div>
</div>

