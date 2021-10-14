import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import { usePostsQuery } from '../graphql/generated'
import styles from '../styles/Home.module.css'

const Home: NextPage = () => {
  const { data, refetch } = usePostsQuery();
  const postsData = data ? data.posts : [];

  return (
    <div className={styles.container}>
      <Head>
        <title>rainbox</title>
        <meta name="description" content="rainbox" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          rainbox
        </h1>

        <div className={styles.description}>
          <ul>
            {postsData.map(post => {
              return <li key={post.id}>{post.tags.join(', ')}</li>
            })}
          </ul>
        </div>
      </main>

      <footer className={styles.footer}>
        <a
          href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          Powered by{' '}
          <span className={styles.logo}>
            <Image src="/vercel.svg" alt="Vercel Logo" width={72} height={16} />
          </span>
        </a>
      </footer>
    </div>
  )
}

export default Home
