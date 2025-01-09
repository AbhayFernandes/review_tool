<script>
	let { children } = $props();
</script>

<header>
	<nav>
		<ul>
			<li><a href="/">Home</a></li>
			<li><a href="/reviews">Reviews</a></li>
			<li><a href="/upload-diff">Upload Diff</a></li>
		</ul>
	</nav>
</header>

<main>
	{@render children()}
</main>
