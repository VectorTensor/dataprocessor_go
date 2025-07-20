package services

import (
	"context"
	"log"
	"main/codegen"

	"main/models"
)

type DataProcessorServer struct {
	codegen.UnimplementedDataprocessorServer
}

func (s *DataProcessorServer) SayHello(ctx context.Context, req *codegen.HelloRequest) (*codegen.HelloReply, error) {

	log.Printf("Receieved %s", req.GetName())

	return &codegen.HelloReply{
		Message: "Hello " + req.GetName(),
	}, nil

}

func (s *DataProcessorServer) GetDBSchema(ctx context.Context, req *codegen.GetDBSchemaRequest) (*codegen.GetDBSchemaResponse, error) {

	db := models.DbOrm{}

	db.Initial()
	table_name := req.Tablename

	columns := db.GetSchema(table_name)

	res := codegen.GetDBSchemaResponse{}

	for _, column := range *columns {

		res.Columns = append(res.Columns, &codegen.ColumnInfo{
			ColumnName: column.ColumnName,
			DataType:   column.DataType,
			IsNullable: column.IsNullable,
		})

	}

	return &res, nil

}
