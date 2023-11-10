<script lang="ts">
	import type { Score } from '$lib/types';
	import Paper, { Title, Content } from '@smui/paper';
	import { onMount } from 'svelte';

	let scores: Score[] = [];
	const getRanks = async () => {
		const res = await fetch('/api/game/rank', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				page: 1,
			}),
		})
			.then((res) => res.json())
			.catch((err) => alert('랭킹불러오기실패 ㅠㅠ'));

		if (res.ok === false) {
			alert('랭킹불러오기실패 ㅠㅠ');
		} else {
			scores = res.scores;
		}
	};

	onMount(() => {
		// getRanks();
		scores = [
			{
				Id: '1',
				Score: 100,
				User1: {
					Id: '1',
					Username: 'user1',
				},
				User2: {
					Id: '2',
					Username: 'user2',
				},
			},
			{
				Id: '2',
				Score: 90,
				User1: {
					Id: '3',
					Username: 'user3',
				},
				User2: {
					Id: '4',
					Username: 'user4',
				},
			},
		];
	});
</script>

<div class="container">
	{#each scores as score, i}
		<div class="paper-container">
			<Paper color="primary" variant="outlined" style="width: 350px;">
				<Title>{i + 1}위 : {score.Score}점</Title>
				<Content>
					{score.User1.Username}<br />
					{score.User2.Username}
				</Content>
			</Paper>
		</div>
	{/each}
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 5px;
		height: 100%;
	}
</style>
