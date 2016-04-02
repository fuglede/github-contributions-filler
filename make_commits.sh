# For each date in the output of make_dates.go, make a random change
# and commit it at that date.
dates=$(go run make_dates.go)
while read -r date; do
	echo "Making a commit at $date"
	echo $RANDOM > trash
	GIT_AUTHOR_DATE=$date GIT_COMMITTER_DATE=$date
	git commit -m "Adding trash to trash" trash > /dev/null
done <<< "$dates"
