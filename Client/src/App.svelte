<script>
  import { onMount } from 'svelte';
  let clock = { isLeapSecond: false, hoursFive: 0, hoursOne: 0, minutesFive: 0, minutesOne: 0 };

  async function getTime() {
    const res = await fetch('http://localhost:8080/time');
    clock = await res.json();
  }

  onMount(() => {
    setInterval(getTime, 1000);
  });
</script>

<main class="panel">
  <div class="clock">  
    <div class="light round {clock.isLeapSecond ? 'yellow' : 'off'}"></div>
    
    <div class="row">
      {#each Array(4) as _, i}
        <div class="light {i < clock.hoursFive ? 'red' : 'off'}"></div>
      {/each}
    </div>

    <div class="row">
      {#each Array(4) as _, i}
        <div class="light {i < clock.hoursOne ? 'red' : 'off'}">1</div>
      {/each}
    </div>

    <div class="row">
      {#each Array(11) as _, i}
        <div class="light-small {i < clock.minutesFive ? ((i+1)%3 === 0 ? 'red' : 'yellow') : 'off'}">5</div>
      {/each}
    </div>

    <div class="row">
      {#each Array(4) as _, i}
        <div class="light {i < clock.minutesOne ? 'yellow' : 'off'}">1</div>
      {/each}
    </div>
  </div>  
</main>

<style>
  .panel { display: flex; justify-content: center; align-items: center; height: 100vh; background: #222; }
  .clock { background: #333; padding: 20px; border-radius: 10px; width: 300px; text-align: center; justify-content: center;}
  .row { display: flex; justify-content: center; gap: 5px; margin-top: 10px; }
  .light { height: 30px; width: 51px; border: 1px solid #000; }
  .light-small { height: 30px; width: 14px; border: 1px solid #000; }
  .round { height: 50px; width: 51px; border-radius: 50%; margin: 0 auto; }
  .yellow { background: yellow; }
  .red { background: red; }
  .off { background: #555; }
</style>