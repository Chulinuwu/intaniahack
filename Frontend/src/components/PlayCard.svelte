<script lang="ts">
    import { colorMap } from '$lib/utils/colorMap';
    import { iconMap } from '$lib/utils/iconMap';
  

    const isMobileDevice = () => {
      return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent);
    };
  
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
  
    let touchStartTime: number | null = null;
    const LONG_PRESS_THRESHOLD = 500; // Milliseconds (adjust as needed)
    let isDragging = false; // Add a flag to track if dragging
  
    function handleTouchStart(event: TouchEvent) {
      touchStartTime = Date.now();
      isDragging = false; // Reset the flag
    }
  
    function handleTouchMove(event: TouchEvent) {
      //If the user moves their finger significantly, consider it dragging, not a long press
      isDragging = true;
      touchStartTime = null;
    }
  
  
    function handleTouchEnd(event: TouchEvent) {
      if (!isDragging && touchStartTime !== null && (Date.now() - touchStartTime) >= LONG_PRESS_THRESHOLD) {
        // It's a long press, prevent default context menu
        event.preventDefault();
        // Optionally, add a visual cue here that the long press was registered,
      }
      touchStartTime = null;
      isDragging = false;
    }
  
    function handleTouchCancel(event: TouchEvent) {
      // Touch cancelled (e.g., interrupted by a system event)
      touchStartTime = null;
      isDragging = false;
    }
  
    function handleContextMenu(event: MouseEvent) {
      // Prevent the default context menu on desktop browsers
      event.preventDefault();
    }
  
    function handleMousedown(event:MouseEvent){  // added the function for support for the desktop
      event.stopPropagation();
    }
  </script>
  
  <div class="card border border-white rounded-lg shadow-lg bg-gray-800 text-white w-[66px] h-[98px]"
      draggable="true"
      on:dragstart
      on:mousedown={handleMousedown}
  >
    <div class={`header p-0.5 ${colorMap[type]} rounded-t-lg flex justify-center items-center`}>
      <img src={iconMap[type]} alt="Money Icon" class="h-2" />
    </div>
    <div class="image-container">
      <img
        src={pic}
        alt={description}
        class="w-[66px] h-[60px] object-cover border-t border-b border-white"
        on:touchstart={handleTouchStart}
        on:touchmove={handleTouchMove}
        on:touchend={handleTouchEnd}
        on:touchcancel={handleTouchCancel}
        on:contextmenu={handleContextMenu}
        style="cursor: pointer; -webkit-user-select: none; -khtml-user-select: none; -moz-user-select: none; -o-user-select: none; user-select: none;"
      />
    </div>
  
    <div class="flex flex-col items-center justify-center text-center">
      <div class="text-[8px] font-medium leading-none">{description}</div>
  
      <div class="flex justify-center items-center gap-0.5">
        {#each stats as stat}
          <div class={`flex items-center justify-center`}>
            <span class={`w-2 h-2 mr-0.5 rounded-full border border-white ${colorMap[stat.name] || 'bg-gray-400'}`} />
            <!-- <img src={colorMap[stat.name]} alt={`${stat.name} icon`} class="w-2 h-2 mr-0.5" /> -->
            <span class="font-light text-[8px]">{stat.value}</span>
          </div>
        {/each}
      </div>
    </div>
  </div>