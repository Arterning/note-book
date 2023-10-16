
## 自增和自减的demo
<script>
	let count = 0;

	$: if (count >= 10) {
		alert(`count is dangerously high!`);
		count = 9;
	}

	$: if (count < 0) {
		alert(`count is low`)
		count = 0
	}

	function handleClick() {
		count += 1;
	}
	function downClick() {
		count -= 1;
	}
</script>

<button on:click={handleClick}>
	Clicked {count}
	{count === 1 ? 'time' : 'times'}
</button>

<button on:click={downClick}>
	Clicked {count}
</button>




