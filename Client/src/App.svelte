<script>
  // Importiere die notwendigen Funktionen von Svelte
  import { onMount } from "svelte";
  import "./App.css";
  let clock = {
    isLeapSecond: false,
    hoursFive: 0,
    hoursOne: 0,
    minutesFive: 0,
    minutesOne: 0,
    secondsFive: 0,
    secondsOne: 0,
    seconds: 0,
  };

  let showDigitalTime = false;

  function toggleTime() {
    showDigitalTime = !showDigitalTime;
  }

  // Asynchrone Funktion um die Zeit vom Backend zu holen
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

<!-- Hauptsektion-->
<main class="mainPanel">
  <div class="clock-assembly">
    <div class="clock-container">
      <header class="clock-header">
        <h1>Berlin-Uhr</h1>
      </header>
      <div class="seconds-orb {clock.isLeapSecond ? 'yellow' : 'off'}">
        {#if showDigitalTime}
          <span class="orb-value">{clock.seconds}</span>
        {/if}
      </div>
      <div class="case-line"></div>

      <!-- Stunden -->
      <div class="row">
        {#each Array(4) as _, i}
          <div class="light {i < clock.hoursFive ? 'red' : 'off'}">
            {#if showDigitalTime && i < clock.hoursFive}
              <span class="segment-value">5h</span>
            {/if}
          </div>
        {/each}
      </div>

      <div class="row">
        {#each Array(4) as _, i}
          <div class="light {i < clock.hoursOne ? 'red' : 'off'}">
            {#if showDigitalTime && i < clock.hoursOne}
              <span class="segment-value">1h</span>
            {/if}
          </div>
        {/each}
      </div>

      <!-- Minuten -->
      <div class="row row-11">
        {#each Array(11) as _, i}
          <div
            class="light-small {i < clock.minutesFive
              ? (i + 1) % 3 === 0
                ? 'red'
                : 'yellow'
              : 'off'}"
          >
            {#if showDigitalTime && i < clock.minutesFive}
              <span class="segment-value">5m</span>
            {/if}
          </div>
        {/each}
      </div>

      <div class="row">
        {#each Array(4) as _, i}
          <div class="light {i < clock.minutesOne ? 'yellow' : 'off'}">
            {#if showDigitalTime && i < clock.minutesOne}
              <span class="segment-value">1m</span>
            {/if}
          </div>
        {/each}
      </div>

      {#if showDigitalTime}
        <div class="case-line"></div>

        <!-- Sekunden  -->
        <div class="row row-11">
          {#each Array(11) as _, i}
            <div
              class="light-small-secsFive {i < clock.secondsFive
                ? 'amber'
                : 'off'}"
            >
              {#if showDigitalTime && i < clock.secondsFive}
                <span class="segment-value">5s</span>
              {/if}
            </div>
          {/each}
        </div>

        <div class="row">
          {#each Array(4) as _, i}
            <div
              class="light-small-secsOne {i < clock.secondsOne
                ? 'amber'
                : 'off'}"
            >
              {#if showDigitalTime && i < clock.secondsOne}
                <span class="segment-value">1s</span>
              {/if}
            </div>
          {/each}
        </div>
      {/if}

      <div class="help-section">
        <button
          class="help-button {showDigitalTime ? 'active' : ''}"
          on:click={toggleTime}
          title="Erklärt die Uhrzeit"
        >
          <span class="help-icon">?</span>
        </button>
        {#if showDigitalTime}
          <div class="digital-time">
            {String(clock.hoursFive * 5 + clock.hoursOne).padStart(
              2,
              "0",
            )}:{String(clock.minutesFive * 5 + clock.minutesOne).padStart(
              2,
              "0",
            )}:{String(clock.seconds).padStart(2, "0")}
          </div>
        {/if}
      </div>
    </div>
  </div>
</main>
