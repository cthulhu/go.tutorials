// 
package main

import fmt "fmt" 

// struct type declaretion
type Point struct { x, y float };

type Point3d struct{
  Point;
  z float;
};

type Point4d struct{
  Point3d;
  float; // have no idea how to call that dimension :)
};


func main( ){
  fmt.Println( "Structs Lesson" );
  
  // simple declaration
  var point struct { x, y float }
  point.x = 10; // simple accessing just like in C
  point.y = 20; // simple accessing just like in C
  // nice trick to see what inside the struct  
  fmt.Println( point );

  // declaration using previously defined type of struct
  var point2 Point;
  fmt.Println( point2 );
  
  // pointer to struct variable has a trick comperatively to C pointers
  var ppoint *Point = new(Point);
  *ppoint = point2;
  fmt.Println( ppoint ); 
  (*ppoint).x = 10; // accessing over the pointer
  ppoint.x = 10;    // the same '.' accessor to fields
  
  // makint struct variables
  
  var point3 Point;
  point3 = Point{ 11, 12 };
  fmt.Println( point3 );   
  
  var point4 Point;
  point4 = Point{ x:11 };
  fmt.Println( point4 );   
  
  var point5 Point;
  point5 = Point{ x:11 };
  fmt.Println( point5 );   
  
  var ppoint2 *Point;
  ppoint2 = &Point{ 11, 12 }; // value will be created automaticaly
  fmt.Println( ppoint2 );   
  
  // exporting structs from other packages
  // struct fields visibility
  var spoint SuperPoint;
  spoint.X = 10;
  spoint.Y = 15;  
  fmt.Println( spoint.X, spoint.Y );
  
  // anonymous struct fields
  var point3d Point3d;
  point3d.x = 10; // those fields are derived
  point3d.y = 20; // from Point struct
  point3d.z = 11;
  fmt.Println( point3d );

  fmt.Println( point3d.Point );   // getting to derived fields
  
  // struct with any types anonymous fields
  point4d := Point4d{ Point3d{ Point{ 10,15 }, 20 }, 30 }; // litterals for anonymous fields should be defined 
  fmt.Println( point4d );    
  
  // methods
  // coomings soon
  
}

// EOF

