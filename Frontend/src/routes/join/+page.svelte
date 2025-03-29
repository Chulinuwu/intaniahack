<script lang="ts">
    import CompetiterCard from "../../components/CompetiterCard.svelte";
    import PlayCard from "../../components/PlayCard.svelte";
    import { colorMap } from '$lib/utils/colorMap';
    import { iconMap } from '$lib/utils/iconMap';
    import { iconMapColor } from '$lib/utils/iconMapColor';
    import TimeLeft from "../../components/TimeLeft.svelte";

    let ageRanges = [
        { label: "0 - 12" },
        { label: "13 - 18" },
        { label: "19 - 22" },
        { label: "23 - 39" },
        { label: "40 - 59" },
        { label: "60 - 79" },
        { label: "80 - 100" }
    ];

    let currentAgeIndex = 1;

    function changeAge(direction: number) {
        currentAgeIndex = Math.max(0, Math.min(ageRanges.length - 1, currentAgeIndex + direction));
    }

    function getDeck() {
        console.log("Deck clicked!");
    }
</script>

<div
    class="flex flex-col gap-2 items-center justify-center"
    style="background-image: url('src/lib/assets/image/play-bg.png'); background-size: cover; background-position: center; height: 100vh;"
>
    <div class="flex gap-5">
        <CompetiterCard
            profileImage="src/lib/assets/image/profile.jpg"
            playerName="John Doe"
            money={1000}
            happiness={80}
            knowledge={90}
            relationship={70}
            age0_12={[
                { type: "money", description: "Learning and growing." },
                { type: "happiness", description: "Exploring and discovering." },
                { type: "relationship", description: "Building a career and relationships." },
                { type: "knowledge", description: "Learning and growing." }, 
            ]}
            age13_18={[
                { type: "money", description: "Learning and growing." },
                { type: "knowledge", description: "Exploring and discovering." },
                { type: "relationship", description: "Building a career and relationships." }
            ]}
        />
        <button class="text-white mt-8" on:click={() => changeAge(-1)}>&#9664;</button>
        <div class="flex flex-col text-white text-center">
            <div class="text-sm">GET THE MOST MONEY</div>
            <TimeLeft />
            <div class="text-xs">AGE 1 - 10</div>
            <div class="flex mt-2 gap-3">
                <PlayCard
                    type="money"
                    pic="src/lib/assets/image/play/money.png"
                    description="Money Card"
                    money={5}
                    happiness={2}
                    knowledge={4}
                    relationship={0}
                />
                <PlayCard
                    type="knowledge"
                    pic="src/lib/assets/image/play/money.png"
                    description="knowledge Card"
                    money={5}
                    happiness={2}
                    knowledge={10}
                    relationship={0}
                />
                <PlayCard
                    type="relationship"
                    pic="src/lib/assets/image/play/money.png"
                    description="relationship Card"
                    money={5}
                    happiness={2}
                    knowledge={10}
                    relationship={0}
                />
                <PlayCard
                    type="happiness"
                    pic="src/lib/assets/image/play/money.png"
                    description="happiness Card"
                    money={5}
                    happiness={2}
                    knowledge={10}
                    relationship={0}
                />
            </div>
        </div>
        <button class="text-white mt-8" on:click={() => changeAge(1)}>&#9654;</button>
        <CompetiterCard
            profileImage="src/lib/assets/image/profile.jpg"
            playerName="John Doe"
            money={1000}
            happiness={80}
            knowledge={90}
            relationship={70}
            age0_12={[
                { type: "money", description: "Learning and growing." },
                { type: "happiness", description: "Exploring and discovering." },
                { type: "relationship", description: "Building a career and relationships." },
                { type: "knowledge", description: "Learning and growing." },
                // { type: "knowledge", description: "Exploring and discovering." },
                // { type: "relationship", description: "Building a career and relationships." },
                // { type: "money", description: "Learning and growing." },
                // { type: "knowledge", description: "Exploring and discovering." },
                
            ]}
            age13_18={[
                { type: "money", description: "Learning and growing." },
                { type: "knowledge", description: "Exploring and discovering." },
                { type: "relationship", description: "Building a career and relationships." }
            ]}
        />
    </div>
    <div class="flex text-white gap-5">
        <button on:click={() => getDeck()} class="w-[166px] h-auto p-0 border-0 bg-transparent">
            <img src="src/lib/assets/image/play/random-deck-button.svg" alt="desk" class="w-full h-auto" />
        </button>
        <div class="flex flex-col gap-3">
            <div class="flex justify-between">
                <div class="flex flex-col">
                    <div class="text-2xl">
                        REMAINING COST
                    </div>
                    <div class="flex gap-2">
                        <div class="flex items-center px-1">
                            <img src={iconMapColor['money']} alt={`money icon`} class="h-4 mr-0.5" />
                            <span class="font-medium text-sm">10</span>
                        </div>
                        <div class="flex items-center px-1">
                            <img src={iconMapColor['happiness']} alt={`happiness icon`} class="h-4 mr-0.5" />
                            <span class="font-medium text-sm">10</span>
                        </div>
                        <div class="flex items-center px-1">
                            <img src={iconMapColor['knowledge']} alt={`knowledge icon`} class="h-4 mr-0.5" />
                            <span class="font-medium text-sm">10</span>
                        </div>
                        <div class="flex items-center px-1">
                            <img src={iconMapColor['relationship']} alt={`relationship icon`} class="h-4 mr-0.5" />
                            <span class="font-medium text-sm">10</span>
                        </div>
                    </div>
                </div>
                <div class="flex flex-col items-center justify-center">
                    <div class="flex items-center gap-2">
                        <img src='src/lib/assets/icon/PlayCard/heart.svg' alt="heart icon" class="h-4 mr-0.5" />
                        <div class="relative w-32 h-4 bg-[#D9D9D9] rounded-full overflow-hidden">
                            <div
                                class="absolute top-0 left-0 h-full bg-[#FF8787]"
                                style="width: calc((1 / 8) * 100%)"
                            ></div>
                        </div>
                        <span class="font-medium text-sm">1</span>
                    </div>
                    <div class="flex items-center gap-2">
                        <img src='src/lib/assets/icon/PlayCard/hourglass.svg' alt="hourglass icon" class="h-4 w-4 mr-0.5" />
                        <div class="relative w-32 h-4 bg-[#D9D9D9] rounded-full overflow-hidden">
                            <div
                                class="absolute top-0 left-0 h-full bg-[#99624E]"
                                style="width: calc((7 / 8) * 100%)"
                            ></div>
                        </div>
                        <span class="font-medium text-sm">7</span>
                    </div>
                </div>
            </div>
            <div class="flex gap-2 w-[530px] h-[114px] bg-[#474848] border border-white rounded-md items-center justify-center">
                        <PlayCard
                            type="happiness"
                            pic="src/lib/assets/image/play/money.png"
                            description="happiness Card"
                            money={5}
                            happiness={2}
                            knowledge={10}
                            relationship={0}
                        />
                        <PlayCard
                            type="happiness"
                            pic="src/lib/assets/image/play/money.png"
                            description="happiness Card"
                            money={5}
                            happiness={2}
                            knowledge={10}
                            relationship={0}
                        />
                        <PlayCard
                            type="happiness"
                            pic="src/lib/assets/image/play/money.png"
                            description="happiness Card"
                            money={5}
                            happiness={2}
                            knowledge={10}
                            relationship={0}
                        />
                        <PlayCard
                            type="happiness"
                            pic="src/lib/assets/image/play/money.png"
                            description="happiness Card"
                            money={5}
                            happiness={2}
                            knowledge={10}
                            relationship={0}
                        />
                        <PlayCard
                            type="happiness"
                            pic="src/lib/assets/image/play/money.png"
                            description="happiness Card"
                            money={5}
                            happiness={2}
                            knowledge={10}
                            relationship={0}
                        />
                        <PlayCard
                            type="happiness"
                            pic="src/lib/assets/image/play/money.png"
                            description="happiness Card"
                            money={5}
                            happiness={2}
                            knowledge={10}
                            relationship={0}
                        />
            </div>
        </div>
        <div class="flex flex-col gap-2 items-center justify-center">
            <div class="flex gap-2">
                <div class="w-[75px] relative">
                    <img src="src/lib/assets/image/play/job-bg.svg" alt="job" class="w-full h-auto" />
                    <div class="absolute inset-0 flex items-center justify-center text-white font-medium text-sm">
                        job
                    </div>
                </div>
                <div class="w-[75px] relative">
                    <img src="src/lib/assets/image/play/investment-bg.svg" alt="investment" class="w-full h-auto" />
                    <div class="absolute inset-0 flex items-center justify-center text-white font-medium text-sm">
                        investment
                    </div>
                </div>
            </div>
            <div class="flex gap-2 relative">
                <div class="w-[75px] relative">
                    <img src="src/lib/assets/image/play/love-bg.svg" alt="love" class="w-full h-auto" />
                    <div class="absolute inset-0 flex items-center justify-center text-white font-medium text-sm">
                        love
                    </div>
                </div>
                <div class="w-[75px] relative">
                    <img src="src/lib/assets/image/play/bin.svg" alt="bin" class="w-full h-[73px]" />
                </div>
            </div>
        </div>
    </div>
</div>