org='openrba'
token='be590e7995f8ecc012a7abb856d13343f99fdcd0'

# function get_repos {
#   curl \
#   -u username:$token \
#   -H "Accept: application/vnd.github.v3+json" \
#   https://api.github.com/orgs/$org/repos | \
#   jq '.[] | .name'
# }

# function get_readme {
#   repo=$1
#   curl \
#   -u username:$token \
#   https://raw.githubusercontent.com/$org/$repo/master/README.md
# }


# repos = $(curl -u username:$token \
#                -H "Accept: application/vnd.github.v3+json" \
#               https://api.github.com/orgs/$org/repos | \
#               jq '.[] | .name')

for repo in $(curl -u username:$token \
               -H "Accept: application/vnd.github.v3+json" \
              https://api.github.com/orgs/$org/repos \
              | grep -oP '(?<=^\s{4}"name": ").*?[^\\](?=",)')
do
  echo "https://raw.githubusercontent.com/$org/$repo/master/README.md"
  curl \
  -u username:$token \
  https://raw.githubusercontent.com/$org/$repo/master/README.md > docs/$repo.md
done
  