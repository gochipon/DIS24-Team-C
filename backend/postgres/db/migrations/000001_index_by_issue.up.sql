-- Create index on comments.issue_url
CREATE INDEX idx_comments_issue_url ON public.comments (issue_url);

-- Create index on issues.number
CREATE INDEX idx_issues_number ON public.issues (number);
