package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// function to get campaign
func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	// if userID is not 0 call the function FindByID
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	// if userID is 0 call the function FindAll
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

// function to get campaign by campaign id
func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByCampaignID(input.CampaignID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

// function to create campaign
func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	stringSlug := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(stringSlug)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}

// function to update campaign
func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByCampaignID(inputID.CampaignID)
	if err != nil {
		return campaign, err
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.GoalAmount = inputData.GoalAmount
	campaign.Perks = inputData.Perks

	newCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}