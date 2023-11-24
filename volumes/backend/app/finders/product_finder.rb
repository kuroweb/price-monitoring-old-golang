# TODO: 各サービス横断のFinderクラスとして作り直す
class ProductFinder
  attr_reader :params

  def initialize(params: {})
    @params = params
  end

  def execute
    products = base_scope
    products = by_id(products)
    products = by_user_id(products)
    products = by_name(products)
    products = by_price(products)

    order(products)
  end

  private

  def base_scope
    Product.all
  end

  def by_id(products)
    return products unless params[:id]

    products.where(id: params[:id])
  end

  def by_user_id(products)
    return products unless params[:user_id]

    products.where(user_id: params[:user_id])
  end

  def by_name(products)
    return products unless params[:name]

    products.where(name: params[:name])
  end

  def by_price(products)
    return products unless params[:price]

    products.where(price: params[:price])
  end

  def order(products)
    return products unless params[:sort]

    products.order_by(params[:sort])
  end
end
