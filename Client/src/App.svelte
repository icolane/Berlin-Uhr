<script>
  import { onMount } from "svelte";
  import "./App.css";

  let clock = {
    isLeapSecond: false,
    hoursFive: 0,
    hoursOne: 0,
    minutesFive: 0,
    minutesOne: 0,
    seconds: 0,
  };

  async function getTime() {
    try {
      const res = await fetch("http://localhost:8080/time");
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
      <h1>Berlin-Uhr</h1>
    </header>
    <div class="light round {clock.isLeapSecond ? 'yellow' : 'off'}"></div>

    <div class="row">
      {#each Array(4) as _, i}
        <div
          class="light {i < clock.hoursFive ? 'red' : 'off'}"
          data-value="5h"
        ></div>
      {/each}
    </div>

    <div class="row">
      {#each Array(4) as _, i}
        <div
          class="light {i < clock.hoursOne ? 'red' : 'off'}"
          data-value="1h"
        ></div>
      {/each}
    </div>

    <!-- Reihe für 5 Minuten-->
    <div class="row">
      {#each Array(11) as _, i}
        <div
          class="light-small {i < clock.minutesFive
            ? (i + 1) % 3 === 0
              ? 'red'
              : 'yellow'
            : 'off'}"
          data-value=""
        ></div>
      {/each}
    </div>

    <div class="row">
      {#each Array(4) as _, i}
        <div
          class="light {i < clock.minutesOne ? 'yellow' : 'off'}"
          data-value=""
        ></div>
      {/each}
    </div>

    <div class="digital-time">
      {String(clock.hoursFive * 5 + clock.hoursOne).padStart(2, "0")}:{String(
        clock.minutesFive * 5 + clock.minutesOne,
      ).padStart(2, "0")}
    </div>
  </div>

  <footer>MENGENLEHREUHR</footer>
</main>
