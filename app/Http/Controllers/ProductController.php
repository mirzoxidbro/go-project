<?php

namespace App\Http\Controllers;

use App\Http\Requests\ProductRequest;
use App\Http\Services\ProductService;
use Illuminate\Http\Request;

class ProductController extends Controller
{
    protected $service;

    public function __construct(){
    $this->service = new ProductService;
}

    public function index(){
        $product = $this->service->index();
        return response()->json([
            'product' => $product
        ]);
    }

    public function store(ProductRequest $request){
        $product = $this->service->store($request->validated());
        return response()->json([
            'product' => $product
        ]);
    }

    public function show($id){
        $product = $this->service->show($id);
        return response()->json([
            'product' => $product
        ]);
    }

    public function update(ProductRequest $request, $id){
        $product = $this->service->update($request->validated(), $id);
        return response()->json([
            'product' => $product
        ]);
    }

    public function destroy($id){
     $product = $this->service->destroy($id);
     return response()->json([
            'message' => 'deleted successfully'
        ]);
    }
}
