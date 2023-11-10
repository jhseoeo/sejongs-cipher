export interface User {
	Id: string;
	Username: string;
}

export interface Score {
	Id: string;
	User1: User;
	User2: User;
	Score: number;
}
