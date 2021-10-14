import { Injectable } from '@nestjs/common';
import { PostsArgs } from './dto/posts.args';
import { Post } from './models/post.model';

@Injectable()
export class PostsService {
  async findOneById(id: string): Promise<Post> {
    return {
      id: 'test',
      createdAt: new Date(),
      tags: ['tag1', 'tag2'],
      thumbnailUrl: 'https://placeimg.com/200/200/tech/0.jpg',
    };
  }

  async findAll(postsArgs: PostsArgs): Promise<Post[]> {
    const posts = [];
    for (let i = 0; i < 100; i++) {
      posts.push({
        id: `test-id-${i}`,
        createdAt: new Date(),
        tags: [`tag${i}`, `tag${i + 1}`],
        thumbnailUrl: `https://placeimg.com/200/200/tech/${i}.jpg`,
      });
    }
    return posts;
  }
}
