package model

import (
	"context"
	"github.com/fadhilthomas/go-repository-audit/config"
	"github.com/jomei/notionapi"
	"github.com/pkg/errors"
)

func OpenNotionDB() (client *notionapi.Client) {
	notionToken := config.GetStr(config.NOTION_TOKEN)
	client = notionapi.NewClient(notionapi.Token(notionToken))
	return client
}

func QueryNotionRepositoryUser(client *notionapi.Client, repositoryName string, userLogin string) (output []notionapi.Page, err error) {
	databaseId := config.GetStr(config.NOTION_DATABASE)

	databaseQueryRequest := &notionapi.DatabaseQueryRequest{
		CompoundFilter: &notionapi.CompoundFilter{
			notionapi.FilterOperatorAND: []notionapi.PropertyFilter{
				{
					Property: "Repository",
					Select: &notionapi.SelectFilterCondition{
						Equals: repositoryName,
					},
				},
				{
					Property: "User Login",
					Select: &notionapi.SelectFilterCondition{
						Equals: userLogin,
					},
				},
			},
		},
	}

	res, err := client.Database.Query(context.Background(), notionapi.DatabaseID(databaseId), databaseQueryRequest)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return res.Results, nil
}

func QueryNotionRepository(client *notionapi.Client, repositoryName string) (output []notionapi.Page, err error) {
	databaseId := config.GetStr(config.NOTION_DATABASE)

	databaseQueryRequest := &notionapi.DatabaseQueryRequest{
		PropertyFilter: &notionapi.PropertyFilter{
			Property: "Repository",
			Select: &notionapi.SelectFilterCondition{
				Equals: repositoryName,
			},
		},
	}

	res, err := client.Database.Query(context.Background(), notionapi.DatabaseID(databaseId), databaseQueryRequest)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return res.Results, nil
}

func InsertNotionRepository(client *notionapi.Client, repository GitHubRepository) (output *notionapi.Page, err error) {
	databaseId := config.GetStr(config.NOTION_DATABASE)

	pageInsertQuery := &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			DatabaseID: notionapi.DatabaseID(databaseId),
		},
		Properties: notionapi.Properties{
			"Organization": notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{
						Text: notionapi.Text{
							Content: repository.OrganizationName,
						},
					},
				},
			},
			"Repository": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: repository.RepositoryName,
				},
			},
			"Owner": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: repository.RepositoryOwner,
				},
			},
			"User Login": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: repository.UserLogin,
				},
			},
			"Admin": notionapi.CheckboxProperty{
				Checkbox: repository.Permission["admin"],
			},
			"Maintain": notionapi.CheckboxProperty{
				Checkbox: repository.Permission["maintain"],
			},
			"Pull": notionapi.CheckboxProperty{
				Checkbox: repository.Permission["pull"],
			},
			"Push": notionapi.CheckboxProperty{
				Checkbox: repository.Permission["push"],
			},
			"Status": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: "open",
				},
			},
		},
	}

	res, err := client.Page.Create(context.Background(), pageInsertQuery)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return res, nil
}

func UpdateNotionRepository(client *notionapi.Client, pageId string, repository GitHubRepository, status string) (output *notionapi.Page, err error) {
	pageUpdateQuery := &notionapi.PageUpdateRequest{
		Properties: notionapi.Properties{
			"Organization": notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{
						Text: notionapi.Text{
							Content: repository.OrganizationName,
						},
					},
				},
			},
			"Repository": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: repository.RepositoryName,
				},
			},
			"Owner": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: repository.RepositoryOwner,
				},
			},
			"User Login": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: repository.UserLogin,
				},
			},
			"Admin": notionapi.CheckboxProperty{
				Checkbox: repository.Permission["admin"],
			},
			"Maintain": notionapi.CheckboxProperty{
				Checkbox: repository.Permission["maintain"],
			},
			"Pull": notionapi.CheckboxProperty{
				Checkbox: repository.Permission["pull"],
			},
			"Push": notionapi.CheckboxProperty{
				Checkbox: repository.Permission["push"],
			},
			"Status": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: status,
				},
			},
		},
	}

	res, err := client.Page.Update(context.Background(), notionapi.PageID(pageId), pageUpdateQuery)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return res, nil
}

func UpdateNotionRepositoryStatus(client *notionapi.Client, pageId string, status string) (output *notionapi.Page, err error) {
	pageUpdateQuery := &notionapi.PageUpdateRequest{
		Properties: notionapi.Properties{
			"Status": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: status,
				},
			},
		},
	}

	res, err := client.Page.Update(context.Background(), notionapi.PageID(pageId), pageUpdateQuery)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return res, nil
}