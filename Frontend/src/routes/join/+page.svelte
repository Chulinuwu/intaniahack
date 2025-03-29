<script lang="ts">
  import CompetiterCard from '../../components/CompetiterCard.svelte';
  import PlayCard from '../../components/PlayCard.svelte';
  import { colorMap } from '$lib/utils/colorMap';
  import { iconMap } from '$lib/utils/iconMap';
  import { iconMapColor } from '$lib/utils/iconMapColor';
  import TimeLeft from '../../components/TimeLeft.svelte';
  import { onMount } from 'svelte';

  let currentAgeIndex = 0;
  $: currentAge = ageRanges[currentAgeIndex];

  function changeAge(direction: number) {
    currentAgeIndex = Math.max(
      0,
      Math.min(ageRanges.length - 1, currentAgeIndex + direction)
    );
  }

  function getDeck() {
    console.log('Deck clicked!');
  }

  // Drag and drop variables
  let draggedCard: any = null;
  let draggedIndex: number | null = null;
  let droppedCards0_12: any[] = Array(5).fill(null);
  let droppedCards13_18: any[] = Array(5).fill(null);
  let droppedCards19_22: any[] = Array(5).fill(null);
  let droppedCards23_39: any[] = Array(5).fill(null);
  let droppedCards40_59: any[] = Array(5).fill(null);
  let droppedCards60_79: any[] = Array(5).fill(null);
  let droppedCards80_100: any[] = Array(5).fill(null);
  let dragSource: 'hand' | 'dropzone' | 'deck' | null = null;
  let dragAgeIndex: number | null = null;
  let draggedCardOffset: { x: number; y: number } = { x: 0, y: 0 };
  let handCards: any[] = [];

  let ageRanges = [
    { label: '0 - 12', data: droppedCards0_12 },
    { label: '13 - 18', data: droppedCards13_18 },
    { label: '19 - 22', data: droppedCards19_22 },
    { label: '23 - 39', data: droppedCards23_39 },
    { label: '40 - 59', data: droppedCards40_59 },
    { label: '60 - 79', data: droppedCards60_79 },
    { label: '80 - 100', data: droppedCards80_100 }
  ];

  function getPosition(event: MouseEvent | DragEvent | TouchEvent) {
    if (event instanceof TouchEvent) {
      return {
        x: event.touches[0].clientX,
        y: event.touches[0].clientY
      };
    } else {
      return {
        x: event.clientX,
        y: event.clientY
      };
    }
  }

  function handleDragStart(
    event: DragEvent | TouchEvent,
    card: any,
    index: number,
    source: 'hand' | 'dropzone' | 'deck'
  ) {
    draggedCard = card;
    draggedIndex = index;
    dragSource = source;
    dragAgeIndex = source === 'dropzone' ? currentAgeIndex : null;

    const target = event.target as HTMLElement;
    const rect = target.getBoundingClientRect();
    const pos = getPosition(event);

    draggedCardOffset = {
      x: pos.x - rect.left,
      y: pos.y - rect.top
    };

    if (event instanceof DragEvent) {
      event.dataTransfer?.setData('text/plain', JSON.stringify({ card, index, source }));
      if (event.dataTransfer) {
        event.dataTransfer.effectAllowed = 'move';
      }
    }

    document.body.classList.add('no-scroll');
  }

  function handleDrop(event: DragEvent | TouchEvent, dropIndex: number) {
    event.preventDefault();
    let card: any = null, index: number | null = null, source: 'hand' | 'dropzone' | 'deck' | null = null;

    try {
      if (event instanceof DragEvent) {
        const data = event.dataTransfer?.getData('text/plain');
        if (data) {
          ({ card, index, source } = JSON.parse(data));
        }
      } else if (event instanceof TouchEvent) {
        card = draggedCard;
        index = draggedIndex;
        source = dragSource;
      }
    } catch (e) {
      console.error("Error parsing data during drop:", e);
      return;
    }

    if (!draggedCard) return;

    if (source === 'hand' && index !== null) {
      if (!currentAge.data[dropIndex]) {
        currentAge.data[dropIndex] = card;
        handCards = handCards.filter((_, i) => i !== index);
      }
    } else if (source === 'dropzone' && index !== null) {
      if (!currentAge.data[dropIndex]) {
        currentAge.data[dropIndex] = card;
        ageRanges[currentAgeIndex].data[index] = null;
      }
    } else if (source === 'deck' && index !== null) {
      if (!currentAge.data[dropIndex]) {
        currentAge.data[dropIndex] = card;
        handCards = handCards.filter((_, i) => i !== index);
      }
    }


    handCards = handCards;
    ageRanges.forEach(age => {
      age.data = [...age.data];
    });
    resetDragState();
  }

  function handleTrashDrop(event: DragEvent | TouchEvent) {
    event.preventDefault();
    if (draggedCard && draggedIndex !== null) {
      if (dragSource === 'hand') {
        handCards = handCards.filter((_, i) => i !== draggedIndex);
      } else if (dragSource === 'dropzone' && dragAgeIndex !== null) {
        ageRanges[dragAgeIndex].data[draggedIndex] = null;
        ageRanges[dragAgeIndex].data = [...ageRanges[dragAgeIndex].data];
      }
      handCards = handCards;
    }
    resetDragState();
  }

  function handleDeckDrop(event: DragEvent | TouchEvent) {
    event.preventDefault();
    try {
      if (dragSource === 'dropzone' && draggedIndex !== null && dragAgeIndex !== null) {
        // Add dragged card back to hand
        handCards = [...handCards, draggedCard];
        ageRanges[dragAgeIndex].data[draggedIndex] = null;
        ageRanges[dragAgeIndex].data = [...ageRanges[dragAgeIndex].data];
        handCards = handCards;
      }
    } catch (e) {
      console.error("Error parsing data during drop:", e);
    }
    resetDragState();
  }

  function handleHandDrop(event: DragEvent | TouchEvent) {
    event.preventDefault();
    try {
      if (dragSource === 'dropzone' && draggedIndex !== null && dragAgeIndex !== null) {
        // Move card from dropzone back to hand
        handCards = [...handCards, draggedCard];
        ageRanges[dragAgeIndex].data[draggedIndex] = null;
        ageRanges[dragAgeIndex].data = [...ageRanges[dragAgeIndex].data];
        handCards = handCards;
      }
    } catch (e) {
      console.error("Error parsing data during drop:", e);
    }
    resetDragState();
  }

  function resetDragState() {
    draggedCard = null;
    draggedIndex = null;
    dragSource = null;
    dragAgeIndex = null;
    draggedCardOffset = { x: 0, y: 0 };
    document.body.classList.remove('no-scroll');
  }

  function handleDragEnd() {
    draggedCard = null;
    document.body.classList.remove('no-scroll');
  }

  let mouseX = 0;
  let mouseY = 0;

  function handleMouseMove(event: MouseEvent | TouchEvent) {
    const pos = getPosition(event);
    mouseX = pos.x;
    mouseY = pos.y;
  }

  onMount(() => {
    document.body.style.overflow = 'hidden';
    handCards = [
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
    ];
  });

  function isMobileDevice(): boolean {
    return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent);
  }

  function handleImageClick(imageUrl: string) {
    if (!isMobileDevice()) {
      downloadImage(imageUrl, imageUrl.substring(imageUrl.lastIndexOf('/') + 1));
    } else {
      // Prevent default long press behavior that allows saving images
      // You can also add additional mobile-specific actions here, like opening in a new tab.
      window.open(imageUrl, '_blank'); // Keeping this to open in a new tab
    }
  }

  function downloadImage(imageUrl: string, filename: string) {
    const a = document.createElement('a');
    a.href = imageUrl;
    a.download = filename;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
  }

</script>

<svelte:window on:mousemove={handleMouseMove} on:touchmove={handleMouseMove} />

<div
  class="flex flex-col gap-0 items-center justify-center h-screen w-screen overflow-hidden"
  style="background-image: url('src/lib/assets/image/play-bg.png'); background-size: cover; background-position: center;"
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
        { type: 'money', description: 'Learning and growing.' },
        { type: 'happiness', description: 'Learning and growing.' },
        { type: 'relationship', description: 'Learning and growing.' },
        { type: 'knowledge', description: 'Learning and growing.' }
      ]}
      age13_18={[
        { type: 'money', description: 'Learning and growing.' },
        { type: 'knowledge', description: 'Learning and growing.' },
        { type: 'relationship', description: 'Learning and growing.' }
      ]}
    />
    <button
      class="text-white mt-10 p-2 rounded-full hover:scale-110 disabled:opacity-50"
      on:click={() => changeAge(-1)}
      disabled={currentAgeIndex === 0}
    >
      ◀
      <!-- ลูกศรซ้าย -->
    </button>
    <div class="flex flex-col text-white text-center">
      <div class="text-sm">GET THE MOST MONEY</div>
      <TimeLeft />
      <div class="text-xs">AGE {currentAge.label}</div>
      <!-- Drop Zones สำหรับช่วงอายุปัจจุบัน -->
      <div class="flex mt-2 gap-3 min-h-[160px]">
        {#each currentAge.data as card, i}
          <div
            class="w-24 h-32 border-2 border-dashed border-white rounded-md flex items-center justify-center"
            role="button"
            on:drop|preventDefault={e => handleDrop(e, i)}
            on:dragover|preventDefault
            on:touchmove|preventDefault
          >
            {#if card}
              <div
                draggable="true"
                on:dragstart={e => handleDragStart(e, card, i, 'dropzone')}
                on:dragend={handleDragEnd}
                tabindex="0"
                role="button"
                style="cursor: grab;"
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
      ▶
      <!-- ลูกศรขวา -->
    </button>
    <CompetiterCard
      profileImage="src/lib/assets/image/profile.jpg"
      playerName="John Doe"
      money={1000}
      happiness={80}
      knowledge={90}
      relationship={70}
      age0_12={[
        { type: 'money', description: 'Learning and growing.' },
        { type: 'happiness', description: 'Learning and growing.' },
        { type: 'relationship', description: 'Learning and growing.' },
        { type: 'knowledge', description: 'Learning and growing.' }
      ]}
      age13_18={[
        { type: 'money', description: 'Learning and growing.' },
        { type: 'knowledge', description: 'Learning and growing.' },
        { type: 'relationship', description: 'Learning and growing.' }
      ]}
    />
  </div>
  <div class="flex text-white gap-5">
    <div
      class="w-[166px] h-auto p-0 border-0 bg-transparent"
      on:drop|preventDefault={handleDeckDrop}
      on:dragover|preventDefault
      on:touchmove|preventDefault
    >
      <img
        src="src/lib/assets/image/play/random-deck-button.svg"
        alt="desk"
        class="w-full h-auto"
        draggable="false"
        on:dragstart={e => e.preventDefault()}
      />
    </div>
    <div class="flex flex-col gap-3">
      <div class="flex justify-between">
        <div class="flex flex-col">
          <div class="text-2xl">REMAINING COST</div>
          <div class="flex gap-2">
            <div class="flex items-center px-1">
              <img
                src={iconMapColor['money']}
                alt={`money icon`}
                class="h-4 mr-0.5"
              />
              <span class="font-medium text-sm">10</span>
            </div>
            <div class="flex items-center px-1">
              <img
                src={iconMapColor['happiness']}
                alt={`happiness icon`}
                class="h-4 mr-0.5"
              />
              <span class="font-medium text-sm">10</span>
            </div>
            <div class="flex items-center px-1">
              <img
                src={iconMapColor['knowledge']}
                alt={`knowledge icon`}
                class="h-4 mr-0.5"
              />
              <span class="font-medium text-sm">10</span>
            </div>
            <div class="flex items-center px-1">
              <img
                src={iconMapColor['relationship']}
                alt={`relationship icon`}
                class="h-4 mr-0.5"
              />
              <span class="font-medium text-sm">10</span>
            </div>
          </div>
        </div>
        <div class="flex flex-col items-center justify-center">
          <div class="flex items-center gap-2">
            <img
              src="src/lib/assets/icon/PlayCard/heart.svg"
              alt="heart icon"
              class="h-4 mr-0.5"
            />
            <div class="relative w-32 h-4 bg-[#D9D9D9] rounded-full overflow-hidden">
              <div
                class="absolute top-0 left-0 h-full bg-[#FF8787]"
                style="width: calc((1 / 8) * 100%)"
              />
            </div>
            <span class="font-medium text-sm">1</span>
          </div>
          <div class="flex items-center gap-2">
            <img
              src="src/lib/assets/icon/PlayCard/hourglass.svg"
              alt="hourglass icon"
              class="h-4 w-4 mr-0.5"
            />
            <div class="relative w-32 h-4 bg-[#D9D9D9] rounded-full overflow-hidden">
              <div
                class="absolute top-0 left-0 h-full bg-[#99624E]"
                style="width: calc((7 / 8) * 100%)"
              />
            </div>
            <span class="font-medium text-sm">7</span>
          </div>
        </div>
      </div>
      <div
        class="flex gap-2 w-[530px] h-[114px] bg-[#474848] border border-white rounded-md items-center justify-center"
        on:drop|preventDefault={handleHandDrop}
        on:dragover|preventDefault
        on:touchmove|preventDefault
      >
        {#each handCards as card, index}
          <div
            draggable="true"
            tabindex="0"
            role="button"
            on:dragstart={e => handleDragStart(e, card, index, 'deck')}
            on:dragend={handleDragEnd}
            style="cursor: grab;"
            aria-grabbed="false"
          >
            <PlayCard {...card} />
          </div>
        {/each}
      </div>
    </div>
    <div class="flex flex-col gap-2 items-center justify-center">
      <div class="flex gap-2">
        <div class="w-[75px] relative">
          <img
            src="src/lib/assets/image/play/job-bg.svg"
            alt="job"
            class="w-full h-auto"
          />
          <div
            class="absolute inset-0 flex items-center justify-center text-white font-medium text-sm"
          >
            job
          </div>
        </div>
        <div class="w-[75px] relative">
          <img
            src="src/lib/assets/image/play/investment-bg.svg"
            alt="investment"
            class="w-full h-auto"
          />
          <div
            class="absolute inset-0 flex items-center justify-center text-white font-medium text-sm"
          >
            investment
          </div>
        </div>
      </div>
      <div class="flex gap-2 relative">
        <div class="w-[75px] relative">
          <img
            src="src/lib/assets/image/play/love-bg.svg"
            alt="love"
            class="w-full h-auto"
          />
          <div
            class="absolute inset-0 flex items-center justify-center text-white font-medium text-sm"
          >
            love
          </div>
        </div>
        <div
          class="w-[75px] relative cursor-pointer"
          on:drop|preventDefault={handleTrashDrop}
          on:dragover|preventDefault
          on:touchmove|preventDefault
          class:active={draggedCard !== null}
        >
          <img
            src="src/lib/assets/image/play/bin.svg"
            alt="bin"
            class="w-full h-[73px] transition-all duration-200"
            class:scale-110={draggedCard !== null}
          />
          {#if draggedCard}
            <div class="absolute inset-0 flex items-center justify-center pointer-events-none"></div>
          {/if}
        </div>
      </div>
    </div>
  </div>
</div>

{#if draggedCard}
  <div
    style="position: fixed;
           top: {mouseY - draggedCardOffset.y}px;
           left: {mouseX - draggedCardOffset.x}px;
           pointer-events: none;
           z-index: 10;
           "
  >
    <PlayCard {...draggedCard} />
  </div>
{/if}

<style>
  /* Apply these styles globally to ensure fullscreen */
  :global(body) {
    margin: 0;
    padding: 0;
    overflow: hidden;
  }

  :global(html) {
    height: 100%;
  }

  img {
    -webkit-user-drag: none;
    -khtml-user-drag: none;
    -moz-user-drag: none;
    -o-user-drag: none;
    user-drag: none;
  }

  .no-scroll {
    overflow: hidden;
  }

  [draggable=true] {
    cursor: grab; /* or use 'grabbing' if you prefer */
  }
</style>