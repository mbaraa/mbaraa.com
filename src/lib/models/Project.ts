export default interface Project {
	publicId?: string;
	name: string;
	description: string;
	startYear: string;
	endYear?: string;
	website?: string;
	sourceCode?: string;
	comingSoon?: boolean;
}
