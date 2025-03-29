<script lang="ts">
    import CompetiterCard from "../../components/CompetiterCard.svelte";
    import PlayCard from "../../components/PlayCard.svelte";
    import { iconMapColor } from '$lib/utils/iconMapColor';
    import TimeLeft from "../../components/TimeLeft.svelte";

    let currentAgeIndex = 0;
    $: currentAge = ageRanges[currentAgeIndex];

    function changeAge(direction: number) {
        currentAgeIndex = Math.max(0, Math.min(ageRanges.length - 1, currentAgeIndex + direction));
    }

    function getDeck() {
        console.log("Deck clicked!");
    }

    // เพิ่มตัวแปรสำหรับระบบ drag and drop
    let dropZones: HTMLDivElement[] = [];
    let draggedCard: any = null;
    let draggedIndex: number | null = null;
    let droppedCards0_12: any[] = Array(5).fill(null);
    let droppedCards13_18: any[] = Array(5).fill(null);
    let droppedCards19_22: any[] = Array(5).fill(null);
    let droppedCards23_39: any[] = Array(5).fill(null);
    let droppedCards40_59: any[] = Array(5).fill(null);
    let droppedCards60_79: any[] = Array(5).fill(null);
    let droppedCards80_100: any[] = Array(5).fill(null);
    let dragSource: 'hand' | 'dropzone' | null = null;
    let dragAgeIndex: number | null = null; // เพิ่มตัวแปรเก็บว่า dragging จากช่วงอายุไหน

    let ageRanges = [
        { label: "0 - 12", data: droppedCards0_12 },
        { label: "13 - 18", data: droppedCards13_18 },
        { label: "19 - 22", data: droppedCards19_22 },
        { label: "23 - 39", data: droppedCards23_39 },
        { label: "40 - 59", data: droppedCards40_59 },
        { label: "60 - 79", data: droppedCards60_79 },
        { label: "80 - 100", data: droppedCards80_100 }
    ];

    let selectedCard: any = null;
    let selectedCardIndex: number | null = null;
    let selectedCardSource: 'hand' | 'dropzone' | null = null;
    let selectedAgeIndex: number | null = null; // สำหรับการ์ดจาก dropzone

    // ฟังก์ชันเลือกการ์ด
    function selectCard(card: any, index: number, source: 'hand' | 'dropzone', ageIndex: number | null = null) {
        // ถ้ากดการ์ดที่เลือกอยู่แล้ว ให้ยกเลิกการเลือก
        if (selectedCardIndex === index && selectedCardSource === source && selectedAgeIndex === ageIndex) {
            selectedCard = null;
            selectedCardIndex = null;
            selectedCardSource = null;
            selectedAgeIndex = null;
        } else {
            selectedCard = card;
            selectedCardIndex = index;
            selectedCardSource = source;
            selectedAgeIndex = ageIndex;
        }
    }

    // ฟังก์ชันย้ายการ์ดไปยังช่องที่เลือก
    function moveCardToSlot(slotIndex: number) {
        if (!selectedCard || selectedCardSource === null) return;

        // ย้ายจาก hand ไปยัง dropzone
        if (selectedCardSource === 'hand' && selectedCardIndex !== null) {
            // ตรวจสอบว่าช่องเป้าหมายว่างไหม
            if (!currentAge.data[slotIndex]) {
                // ย้ายการ์ด
                handCards = handCards.filter((_, i) => i !== selectedCardIndex);
                currentAge.data[slotIndex] = selectedCard;
            }
        }
        // ย้ายระหว่าง dropzone
        else if (selectedCardSource === 'dropzone' && selectedCardIndex !== null && selectedAgeIndex !== null) {
            // ตรวจสอบว่าช่องเป้าหมายว่างไหม
            if (!currentAge.data[slotIndex]) {
                // ย้ายการ์ด
                ageRanges[selectedAgeIndex].data[selectedCardIndex] = null;
                currentAge.data[slotIndex] = selectedCard;
            }
        }

        // รีเซ็ตการเลือก
        resetSelection();
    }

    // ฟังก์ชันทิ้งการ์ด
    function trashCard() {
        if (!selectedCard || selectedCardSource === null) return;

        // ถ้ามาจาก hand
        if (selectedCardSource === 'hand' && selectedCardIndex !== null) {
            handCards = handCards.filter((_, i) => i !== selectedCardIndex);
        }
        // ถ้ามาจาก dropzone
        else if (selectedCardSource === 'dropzone' && selectedCardIndex !== null && selectedAgeIndex !== null) {
            ageRanges[selectedAgeIndex].data[selectedCardIndex] = null;
        }

        // รีเซ็ตการเลือก
        resetSelection();
    }

    // รีเซ็ตการเลือกการ์ด
    function resetSelection() {
        selectedCard = null;
        selectedCardIndex = null;
        selectedCardSource = null;
        selectedAgeIndex = null;
        
        // อัปเดต UI
        handCards = handCards;
        ageRanges.forEach(age => {
            age.data = [...age.data];
        });
    }


    let handCards = [
        {
            type: "happiness",
            pic: "src/lib/assets/image/play/money.png",
            description: "happiness Card",
            money: 5,
            happiness: 2,
            knowledge: 10,
            relationship: 0
        },
        {
            type: "happiness",
            pic: "src/lib/assets/image/play/money.png",
            description: "happiness Card",
            money: 5,
            happiness: 2,
            knowledge: 10,
            relationship: 0
        },
        {
            type: "happiness",
            pic: "src/lib/assets/image/play/money.png",
            description: "happiness Card",
            money: 5,
            happiness: 2,
            knowledge: 10,
            relationship: 0
        },
        {
            type: "happiness",
            pic: "src/lib/assets/image/play/money.png",
            description: "happiness Card",
            money: 5,
            happiness: 2,
            knowledge: 10,
            relationship: 0
        },
        {
            type: "happiness",
            pic: "src/lib/assets/image/play/money.png",
            description: "happiness Card",
            money: 5,
            happiness: 2,
            knowledge: 10,
            relationship: 0
        },
        {
            type: "happiness",
            pic: "src/lib/assets/image/play/money.png",
            description: "happiness Card",
            money: 5,
            happiness: 2,
            knowledge: 10,
            relationship: 0
        }
    ]
</script>

<div
    class="flex flex-col gap-0 items-center justify-center"
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
        <button 
            class="text-white mt-10 p-2 rounded-full hover:scale-110 disabled:opacity-50" 
            on:click={() => changeAge(-1)} 
            disabled={currentAgeIndex === 0}
        >
            &#9664; <!-- ลูกศรซ้าย -->
        </button>
        <div class="flex flex-col text-white text-center">
            <div class="text-sm">GET THE MOST MONEY</div>
            <TimeLeft />
            <div class="text-xs">AGE {currentAge.label}</div>
            <!-- Drop Zones สำหรับช่วงอายุปัจจุบัน -->
            <div class="flex mt-2 gap-3 min-h-[160px]">
                {#each currentAge.data as card, i}
        <div 
            class="w-24 h-32 border-2 border-dashed border-white rounded-md flex items-center justify-center cursor-pointer"
            on:click={() => {
                if (card) {
                    // ถ้ามีการ์ดอยู่แล้ว ให้เลือกการ์ดนั้น
                    selectCard(card, i, 'dropzone', currentAgeIndex);
                } else if (selectedCard) {
                    // ถ้าไม่มีการ์ดแต่มีการ์ดเลือกอยู่ ให้ย้ายการ์ดมาที่นี่
                    moveCardToSlot(i);
                }
            }}
            class:bg-gray-700={selectedCard && !card}
        >
            {#if card}
                <div
                    on:click|stopPropagation={() => selectCard(card, i, 'dropzone', currentAgeIndex)}
                    class:scale-110={selectedCardIndex === i && selectedCardSource === 'dropzone' && selectedAgeIndex === currentAgeIndex}
                    class="transition-transform duration-200"
                >
                    <PlayCard {...card} />
                </div>
            {/if}
        </div>
    {/each}
            </div>
        </div>
        <button 
            class="text-white mt-10 p-2 rounded-full hover:scale-110 disabled:opacity-50" 
            on:click={() => changeAge(1)} 
            disabled={currentAgeIndex === ageRanges.length - 1}
        >
            &#9654; <!-- ลูกศรขวา -->
        </button>
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
                {#each handCards as card, index}
                    <div 
                        on:click={() => selectCard(card, index, 'hand')}
                        class:scale-110={selectedCardIndex === index && selectedCardSource === 'hand'}
                        class="transition-transform duration-200 cursor-pointer"
                    >
                        <PlayCard {...card} />
                    </div>
                {/each}
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
                <div 
                class="w-[75px] relative cursor-pointer transition-transform duration-200"
                on:click={trashCard}
                class:scale-110={selectedCard !== null}
            >
                <img 
                    src="src/lib/assets/image/play/bin.svg" 
                    alt="bin" 
                    class="w-full h-[73px]"
                />
                {#if selectedCard}
                    <div class="absolute inset-0 flex items-center justify-center pointer-events-none">
                    </div>
                {/if}
            </div>
            </div>
        </div>
    </div>
</div>