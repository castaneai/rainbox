import { NotFoundException } from '@nestjs/common';
import { Args, Query, Resolver } from '@nestjs/graphql';
import { PostsArgs } from './dto/posts.args';
import { Post } from './models/post.model';
import { PostsService } from './posts.service';

@Resolver((of) => Post)
export class PostsResolver {
  constructor(private readonly postsService: PostsService) {}

  @Query((returns) => Post)
  async post(@Args('id') id: string): Promise<Post> {
    const post = await this.postsService.findOneById(id);
    if (!post) {
      throw new NotFoundException(id);
    }
    return post;
  }

  @Query((returns) => [Post])
  posts(@Args() postsArgs: PostsArgs): Promise<Post[]> {
    return this.postsService.findAll(postsArgs);
  }
}
