---
description: Review staged changes and commit with a generated message
agent: build
---

Review only the changes that are already staged.

Staged files:

!`git diff --cached --stat`

Staged diff:

!`git diff --cached`

Instructions:

1. If the repository is not initialized or nothing is staged, say so and stop. Never run `git add`.
2. Write one Conventional Commit message: `type(scope): summary`. Add a short body only if the change needs explanation.
3. Show the message and the list of staged files.
4. Ask for confirmation.
5. Only after confirmation, commit the staged files with `git commit -m "<subject>" -m "<body>"`.
6. Do not amend, push, or touch unstaged changes.
