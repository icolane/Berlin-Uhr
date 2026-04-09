<script>
  import { onMount } from 'svelte';
  import './App.css';
  
  let clock = { isLeapSecond: false, hoursFive: 0, hoursOne: 0, minutesFive: 0, minutesOne: 0 };

  async function getTime() {
    try {
      const res = await fetch('http://localhost:8080/time');
      if (res.ok) {
        clock = await res.json();
      }
    } catch (e) {
      console.error("Failed to fetch time", e);
    }
  }

  onMount(() => {
    getTime();
    const interval = setInterval(getTime, 1000);
    return () => clearInterval(interval);
  });
</script>

<main class="panel">
  <div class="clock-container">
    <header class="clock-header">
      <h1>Berlin Clock</h1>
    </header>

    <div class="light round {clock.isLeapSecond ? 'yellow' : 'off'}"></div>
    
    <!-- 5 Hours -->
    <div class="row">
      {#each Array(4) as _, i}
        <div class="light {i < clock.hoursFive ? 'red' : 'off'}" data-value="5h"></div>
      {/each}
    </div>

    <!-- 1 Hour -->
    <div class="row">
      {#each Array(4) as _, i}
        <div class="light {i < clock.hoursOne ? 'red' : 'off'}" data-value="1h"></div>
      {/each}
    </div>

    <!-- 5 Minutes -->
    <div class="row">
      {#each Array(11) as _, i}
        <div class="light-small {i < clock.minutesFive ? ((i+1)%3 === 0 ? 'red' : 'yellow') : 'off'}" data-value="5m"></div>
      {/each}
    </div>

    <!-- 1 Minute -->
    <div class="row">
      {#each Array(4) as _, i}
        <div class="light {i < clock.minutesOne ? 'yellow' : 'off'}" data-value="1m"></div>
      {/each}
    </div>
  </div>

  <footer>MENGENLEHREUHR</footer>
</main>