package userdb

import (
	"fmt"
	"strings"
)

type UserLookup struct {
	ID       *int64
	Username *string
}

func (udb *UserDB) GetUser(lookup UserLookup, includeProfile bool, includeMetrics bool) (*User, *UserProfile, *UserPublicMetrics, error) {
	query := "SELECT u.id, u.name, u.username, u.pfp_url, u.protected"
	joins := []string{}
	scanArgs := []interface{}{}
	var user User
	var profile *UserProfile
	var metrics *UserPublicMetrics

	scanArgs = append(scanArgs, &user.ID, &user.Name, &user.Username, &user.PfpURL, &user.Protected)

	if includeProfile {
		query += ", p.bio, p.website, p.location, p.birthday, p.pinned_post_id"
		joins = append(joins, "LEFT JOIN user_profiles p ON u.id = p.user_id")
		profile = &UserProfile{}
		scanArgs = append(scanArgs, &profile.Bio, &profile.Website, &profile.Location, &profile.Birthday, &profile.PinnedPostID)
	}

	if includeMetrics {
		query += ", m.followers_count, m.following_count, m.post_count, m.listed_count"
		joins = append(joins, "LEFT JOIN user_public_metrics m ON u.id = m.user_id")
		metrics = &UserPublicMetrics{}
		scanArgs = append(scanArgs, &metrics.FollowersCount, &metrics.FollowingCount, &metrics.PostCount, &metrics.ListedCount)
	}

	query += " FROM users u " + strings.Join(joins, " ") + " WHERE "

	var args []interface{}
	if lookup.ID != nil {
		query += "u.id = $1"
		args = append(args, *lookup.ID)
	} else if lookup.Username != nil {
		query += "u.username = $1"
		args = append(args, *lookup.Username)
	} else {
		return nil, nil, nil, fmt.Errorf("either ID or Username must be provided")
	}

	err := udb.db.QueryRow(query, args...).Scan(scanArgs...)
	if err != nil {
		return nil, nil, nil, err
	}

	return &user, profile, metrics, nil
}
