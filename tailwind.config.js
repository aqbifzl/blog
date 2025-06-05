module.exports = {
  content: ["./components/*.templ"],
  theme: {
    extend: {
      typography: () => ({
        green: {
          css: {
            '--tw-prose-body': 'var(--color-black)',
            '--tw-prose-headings': 'var(--color-purple-900)',
            '--tw-prose-lead': 'var(--color-gray-700)',
            '--tw-prose-links': 'var(--color-purple-900)',
            '--tw-prose-bold': 'var(--color-purple-900)',
            '--tw-prose-counters': 'var(--color-purple-600)',
            '--tw-prose-bullets': 'var(--color-purple-400)',
            '--tw-prose-hr': 'var(--color-purple-300)',
            '--tw-prose-quotes': 'var(--color-purple-900)',
            '--tw-prose-quote-borders': 'var(--color-purple-300)',
            '--tw-prose-captions': 'var(--color-purple-700)',
            '--tw-prose-code': 'var(--color-purple-900)',
            '--tw-prose-pre-code': 'var(--color-purple-100)',
            '--tw-prose-pre-bg': 'var(--color-purple-900)',
            '--tw-prose-th-borders': 'var(--color-purple-300)',
            '--tw-prose-td-borders': 'var(--color-purple-200)',
          },
        },
      }),
    },
  },
  plugins: [require("@tailwindcss/typography")],
};

