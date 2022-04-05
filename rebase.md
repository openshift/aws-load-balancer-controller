# Rebasing Procedure

This fork needs to be periodically synced with upstream changes. To do this we
need to run the following steps:

1. Ensure the *upstream*, *downstream* and your own fork of the repository are  
   configured correctly as remotes.
   ```bash
   git remote -v
   # if upstream or downstream are not present then add them
   git remote add upstream https://github.com/kubernetes-sigs/aws-load-balancer-controller
   git remote add downstream https://github.com/openshift/aws-load-balancer-controller
   ```

2. Ensure that the `main` branch is up-to-date with the *downstream* `main`
   branch.
   ```bash
   git switch main
   git pull downstream main
   ```

3. Update the remotes
   ```bash
   git remote update
   ```

4. Checkout the commit from upstream to which you want to rebase the `main`
   branch.
   ```bash
   git checkout $TAG_OR_COMMIT_ID
   git switch -c rebase
   ```

5. Merge the `main` branch into the current branch with
   the [ours](https://git-scm.com/docs/merge-strategies#_merge_strategies)
   merge strategy.
   ```bash
   git merge -s ours main
   ```
   _Note:_ This strategy should not be confused with the `ours` option for 
   the `ort` merge strategy.

7. Determine the changes present in the `main` branch which need to be
   cherry-picked.
   ```bash
   git log --oneline --no-merges $TAG_OR_COMMIT_ID..main
   ```
   
## Message format for cherry-picked commits

The commits which are cherry-picked into the rebase branch should have the 
following format:

* `UPSTREAM: <carry>: $MESSAGE`
  A persistent carry that should probably be picked for the subsequent rebase
  branch. In general, these commits are used to modify behavior from upstream.

* `UPSTREAM: <drop>: $MESSAGE`
  A carry that should probably not be picked for the subsequent rebase branch.
  In general, these commits are used to maintain the codebase in ways that are
  branch-specific, like the update of generated files or dependencies.

* `UPSTREAM: 77870: $MESSAGE`
  The number identifies a PR in the upstream repository(
  i.e. https://github.com/kubernetes-sigs/aws-load-balancer-controller/pull/<pr_id>)
  A commit with this message should only be picked into the subsequent rebase
  branch if the commits of the referenced PR are not included in the upstream
  branch. To check if a given commit is included in the upstream branch, open
  the referenced upstream PR and check any of its commits for the release tag (
  e.g. v.1.20.0) targeted by the new rebase branch.
   

