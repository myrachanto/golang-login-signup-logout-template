package support

import (
	
	s "github.com/myrachanto/asearch/models"
)
var operator = map[string]string{"equal_to":"=","not_equal_to":"<>","less_than":"<",
"greater_than":">","less_than_or_equal_to":">=","greater_than_ro_equal_to":">=",
"in":"in","like":"like","between":"between","not_in":"not_in"}
func SearchData( column, direction,search_column,search_operator, operator, search_query_1, search_query_2, per_page string)(result s.Customer){

	return s.GormDB.Order(column, direction).buildQuery(
		ser.Search_column,
		ser.Search_operator,
		).Scan(&result)
	//.where(s.GormDB)
		// // chcek if search s.GormDB is empty
		// if(search_query_1){ 
		// 	return buildQuery(
		// 	search_column,
		// 	search_operator,
		// 	)
		// }

}

func buildQuery(column, operator, search_query_1, search_query_2, string){
	switch(operator){
		case "equal_to":
		case "not_equal_to":
		case "less_than" :
		case "greater_than":
		case "less_than_or_equal_to":
		case "greater_than_ro_equal_to":
		s.GormDB.where(column, operator, search_s.GormDB_1);
		break;
		case "in":
		s.GormDB.where(column, operator, []string{search_query_1, search_query_2});
		case "not_in":
		s.GormDB.where(column, operator, []string{search_query_1, search_query_2});
		break;
		case "like":
		s.GormDB.where(column, operator, "%"+search_query_1+"%");
		break;
		case "between":
			s.GormDB.where(column, operator, []string{search_query_1, search_query_2});
		   break;
		default:
		return fmt.Println("Invalid Search Operator");
		break;
	}
	//return results, err;
}