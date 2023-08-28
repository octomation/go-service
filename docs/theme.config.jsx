export default {
  project: {
    link: 'https://github.com/octomation/go-service',
  },

  docsRepositoryBase: 'https://github.com/octomation/go-service/blob/main/docs',
  feedback: {
    useLink() {
      return 'https://github.com/octomation/go-service/discussions/new/choose'
    },
  },
  useNextSeoProps() {
    return {
      titleTemplate: '%s',
    }
  },

  head: (
    <>
      <meta charSet="utf-8"/>
      <meta name="viewport" content="width=device-width, initial-scale=1.0"/>

      <meta name="twitter:image:src" content="https://cdn.octolab.org/repo/go-service.png"/>
      <meta name="twitter:site" content="@github"/>
      <meta name="twitter:card" content="summary_large_image"/>
      <meta name="twitter:title" content="Tool"/>
      <meta name="twitter:description" content="🧩 Template for a typical service written on Go."/>
      <meta property="og:image" content="https://cdn.octolab.org/repo/go-service.png"/>
      <meta property="og:image:alt" content="🧩 Service"/>
      <meta property="og:site_name" content="GitHub"/>
      <meta property="og:type" content="object"/>
      <meta property="og:title" content="Service"/>
      <meta property="og:url" content="https://go-service.octolab.org"/>
      <meta property="og:description" content="🧩 Template for a typical service written on Go."/>

      <style>{`
        main a img { display: inline; } /* badges */
      `}</style>
    </>
  ),
  logo: (
    <>
      <img width={24} height={24} src="https://cdn.octolab.org/geek/octolab.png" alt="OctoLab"/>
      <span>Tool</span>
    </>
  ),
  banner: {
    text: <a href="https://github.com/octomation/go-service/releases/tag/v1.0.0" target="_blank">
      🎉 Service v1.0.0 is released. Read more →
    </a>,
  },
  footer: {
    text: <span>
      MIT {new Date().getFullYear()} © <a href="https://github.com/octolab" target="_blank">OctoLab</a>.
    </span>,
  },
}
