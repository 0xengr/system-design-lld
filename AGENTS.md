# AGENTS.md

## Repo shape

- This is a teaching-content repo for low-level system design, not an application. Lessons are Markdown; the TypeScript in them is illustrative and is not a runnable project.
- There is no `package.json`, build, test, lint, or formatter. Do not add tooling or try to execute snippets unless the user explicitly asks.

## Adding or editing lessons

- Top-level directories are numbered chapters (`01-system-design-fundamentals/`); lesson files use `NN-kebab-case-title.md` within a chapter. Follow that layout for new lessons.
- Mirror the existing lesson structure: an H1 title, learning objectives, numbered `##` sections separated by `---`, Mermaid diagrams, tables for comparisons, TypeScript examples, key takeaways and terminology, a practical exercise, and interview questions.
- Lesson code is deliberately simplified. When an example omits production concerns (async failures, retries, idempotency, money in integer minor units), state the caveat in prose rather than silently "fixing" it.

## Commits

- The `/commit` command in `.opencode/commands/commit.md` commits only already-staged changes, never runs `git add`, and uses Conventional Commit messages (`type(scope): summary`).
