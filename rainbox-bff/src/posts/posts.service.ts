import { Injectable } from '@nestjs/common';
import { PostsArgs } from './dto/posts.args';
import { Post } from './models/post.model';

@Injectable()
export class PostsService {
  async findOneById(id: string): Promise<Post> {
    return { id: 'test', createdAt: new Date(), tags: ['tag1', 'tag2'] };
  }

  async findAll(postsArgs: PostsArgs): Promise<Post[]> {
    return [];
  }
}
