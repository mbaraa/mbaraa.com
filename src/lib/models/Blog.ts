export default interface Blog {
  publicId?: string;
  name: string;
  description: string;
  content: string;
  createdAt: Date;
  updatedAt: Date;
}
