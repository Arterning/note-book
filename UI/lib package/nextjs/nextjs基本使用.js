/**
 * 定义 pages/posts/[slug].js 文件
 * 
 * getStaticPaths define route, bind route to component
 * 
 * getStaticProps define the props send to component
 */

export async function getStaticPaths() {
    const slugs = await getSlugs();
    return {
      paths: slugs.map((slug) => ({
        params: { slug },
      })),
      fallback: false,
    };
  }
  
  export async function getStaticProps({ params: { slug } }) {
    console.log('[PostPage] getStaticProps():', slug);
    const post = await getPost(slug);
    return {
      props: { post },
    };
  }
  
  function PostPage({ post }) {
    console.log('[PostPage] render:', post);
    return (
      <>
        <Head>
          <title>{`${post.title} - My Blog`}</title>
        </Head>
        <main>
          <p>{post.date}</p>
          <h1>{post.title}</h1>
          <article dangerouslySetInnerHTML={{ __html: post.body }} />
        </main>
      </>
    );
  }
  
  export default PostPage;
  