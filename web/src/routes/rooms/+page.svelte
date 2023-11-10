<script lang="ts">
	import Card, { Content, Actions } from '@smui/card';
	import Button, { Label } from '@smui/button';
	import type { Room } from '$lib/types';
	import { config } from '$lib/config';
	import { checkLogin } from '$lib/checkLogin';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	let rooms: Room[] = [];
	const getRooms = async () => {
		const res = await fetch(config.apiHost + '/api/room', {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				Authorization: 'Bearer ' + localStorage.getItem('token') ?? '',
			},
		})
			.then((res) => res.json())
			.catch((err) => alert('방목록불러오기실패 ㅠㅠ'));

		if (res.ok === false) {
			alert('방목록불러오기실패 ㅠㅠ');
		} else {
			rooms = res.data.rooms;
		}
	};

	onMount(() => {
		checkLogin().then((logged_in) => {
			if (logged_in === false) {
				alert('로그인이 필요합니다.');
				location.href = '/login';
			}
			// getRooms();
		});

		rooms = [
			{
				Id: '1',
				RoomName: '방이름',
				User1: {
					Id: '1',
					Username: '유저1',
				},
				User2: {
					Id: '2',
					Username: '유저2',
				},
			},
			{
				Id: '2',
				RoomName: '방이름2',
				User1: {
					Id: '3',
					Username: '유저3',
				},
			},
		];
	});
</script>

<div class="container">
	{#each rooms as room, i}
		{@const numUsers = room.User2 ? 2 : 1}
		<Card style="width: 300px;">
			<Content>
				{room.RoomName} ({numUsers}/2)<br />
			</Content>
			<Actions style="display: flex; flex-direction: row-reverse">
				<Button
					on:click={() => {
						goto('/rooms/' + room.Id);
					}}
				>
					<Label>참가</Label>
				</Button>
			</Actions>
		</Card>
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
