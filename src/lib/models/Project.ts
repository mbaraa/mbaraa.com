export default interface Project {
    name: string;
    description: string;
    startYear: string;
    endYear?: string;
    website?: string;
    sourceCode?: string;
    imagePath: string;
    comingSoon?: boolean;
}
