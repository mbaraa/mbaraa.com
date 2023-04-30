export default interface Project {
	publicId?: string;
	name: string;
	description: string;
	startYear: Date;
	endYear?: Date;
	website?: string;
	sourceCode?: string;
	comingSoon?: boolean;
}
