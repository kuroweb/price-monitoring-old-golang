class MercariCrawlSettingRequiredKeyword < ApplicationRecord
  ## associations ##
  belongs_to :mercari_crawl_setting

  ## validations ##
  validates :keyword, presence: true

  ## scopes ##
  ## methods ##
end
