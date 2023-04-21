export default interface Blog {
	publicId?: string;
	name: string;
	description: string;
	readTimes: number;
	likes: number;
	content: string;
	createdAt: Date;
	updatedAt: Date;
}
