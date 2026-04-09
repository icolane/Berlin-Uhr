<script>
  export let clock;
</script>

<div class="clock">
  <div class="light round {clock.isLeapSecond ? 'yellow' : 'off'}"></div>

  <div class="row">
    {#each Array(4) as _, i}
      <div class="light {i < clock.hoursFive ? 'red' : 'off'}"></div>
    {/each}
  </div>

  <div class="row">
    {#each Array(4) as _, i}
      <div class="light {i < clock.hoursOne ? 'red' : 'off'}"></div>
    {/each}
  </div>

  <div class="row">
    {#each Array(11) as _, i}
      <div class="light-small {i < clock.minutesFive ? ((i+1)%3 === 0 ? 'red' : 'yellow') : 'off'}"></div>
    {/each}
  </div>

  <div class="row">
    {#each Array(4) as _, i}
      <div class="light {i < clock.minutesOne ? 'yellow' : 'off'}"></div>
    {/each}
  </div>
</div>

<style>
  .clock {
    background: linear-gradient(135deg, #2c3e50, #34495e);
    padding: 30px;
    border-radius: 20px;
    width: 350px;
    margin: 0 auto;
    box-shadow: 0 10px 30px rgba(0,0,0,0.3);
    text-align: center;
  }
  .row {
    display: flex;
    justify-content: center;
    gap: 8px;
    margin: 15px 0;
  }
  .light {
    height: 40px;
    width: 60px;
    border: 2px solid #000;
    border-radius: 5px;
    transition: all 0.3s ease;
  }
  .light-small {
    height: 40px;
    width: 20px;
    border: 2px solid #000;
    border-radius: 5px;
    transition: all 0.3s ease;
  }
  .round {
    height: 60px;
    width: 60px;
    border-radius: 50%;
    margin: 0 auto 20px;
    border: 2px solid #000;
    transition: all 0.3s ease;
  }
  .yellow {
    background: #f1c40f;
    box-shadow: 0 0 20px #f1c40f;
  }
  .red {
    background: #e74c3c;
    box-shadow: 0 0 20px #e74c3c;
  }
  .off {
    background: #7f8c8d;
    box-shadow: none;
  }
</style>