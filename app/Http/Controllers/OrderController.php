<?php

namespace App\Http\Controllers;

use App\Http\Requests\OrderRequest;
use App\Http\Services\OrderService;
use Illuminate\Http\Request;

class OrderController extends Controller
{
    protected $service;

    public function __construct(){
    $this->service = new OrderService;
}

    public function index(){
        $order = $this->service->index();
        return response()->json([
            'order' => $order
        ]);
    }

    public function store(OrderRequest $request){
        $order = $this->service->store($request->validated());
        return response()->json([
            'order' => $order
        ]);
    }

    public function show($id){
        $order = $this->service->show($id);
        return response()->json([
            'order' => $order
        ]);
    }


    public function update(OrderRequest $request, $id){
        $order = $this->service->update($request->validated(), $id);
        return response()->json([
            'order' => $order
        ]);
    }

    public function destroy($id){
     $order = $this->service->destroy($id);
     return response()->json([
            'message' => 'deleted successfully'
        ]);
    }
}
