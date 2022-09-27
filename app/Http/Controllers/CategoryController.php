<?php

namespace App\Http\Controllers;

use App\Http\Requests\CategoryRequest;
use App\Http\Services\CategoryService;
use Illuminate\Http\Request;

class CategoryController extends Controller{

    protected $service;

    public function __construct(){
    $this->service = new CategoryService;
}

    public function index(){
        $category = $this->service->index();
        return response()->json([
            'category' => $category
        ]);
    }

    public function store(CategoryRequest $request){
        $category = $this->service->store($request->validated());
        return response()->json([
            'category' => $category
        ]);
    }

    public function show($id){
        $category = $this->service->show($id);
        return response()->json([
            'category' => $category
        ]);
    }

    public function update(CategoryRequest $request, $id){
        $category = $this->service->update($request->validated(), $id);
        return response()->json([
            'category' => $category
        ]);
    }

    public function destroy($id){
     $category = $this->service->destroy($id);
     return response()->json([
            'message' => 'deleted successfully'
        ]);
    }
}
